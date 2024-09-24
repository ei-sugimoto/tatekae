import { Flex, Heading, Spacer, Text } from '@chakra-ui/react';
import { useParams } from 'react-router-dom';
import {
  BillServiceListResponse,
  BillServiceSumarizeByProjectResponse,
} from '../gen/proto_bill/v1/bill_pb';
import { ConnectError } from '@connectrpc/connect';
import { Client } from '../use-client';
import { BillService } from '../gen/proto_bill/v1/bill_connect';
import { useEffect, useState } from 'react';
import { ProjectServiceGetResponse } from '../gen/proto_project/v1/project_pb';
import { ProjectService } from '../gen/proto_project/v1/project_connect';
import BillForm from '../components/billForm';

export function ProjectDetail() {
  const { id } = useParams<{ id: string }>();
  const [billList, setBillList] = useState<BillServiceListResponse | null>(
    null
  );
  const [error, setError] = useState<ConnectError | null>(null);

  const [project, setProject] = useState<ProjectServiceGetResponse | null>(
    null
  );

  const [sumarize, setSumarize] =
    useState<BillServiceSumarizeByProjectResponse | null>(null);

  const [memberList, setMemberList] = useState<Array<{
    id: number;
    name: string;
  }> | null>(null);
  useEffect(() => {
    if (!id) {
      setError(new ConnectError('id is not provided'));
      return;
    }

    fetch(id);
  }, [id]);

  if (error) {
    return (
      <>
        <Heading>Project Detail</Heading>
        <Spacer mt={5} />
        <Text fontSize={'2xl'}>{error.message}</Text>
      </>
    );
  }

  const fetch = (id: string) => {
    getBillList(parseInt(id)).then((res) => {
      if (res instanceof ConnectError) {
        setError(res);
      } else {
        setBillList(res);
      }
    });

    getProjectDetail(parseInt(id)).then((res) => {
      if (res instanceof ConnectError) {
        setError(res);
      } else {
        setProject(res);
      }
    });

    getSumerize(parseInt(id)).then((res) => {
      if (res instanceof ConnectError) {
        setError(res);
      } else {
        setSumarize(res);
      }
    });

    getMemberList(parseInt(id)).then((res) => {
      if (res instanceof ConnectError) {
        setError(res);
      } else {
        const response = res.members.map((member) => {
          return { id: member.id, name: member.name };
        });
        setMemberList(response);
      }
    });
  };

  return (
    <>
      <Flex direction={'column'} align={'center'}>
        <Heading>{project?.name}</Heading>
        <Spacer mt={5} />
        <Text fontSize={'2xl'}>Member list</Text>
        {memberList &&
          memberList.map((user) => {
            return (
              <>
                <Flex justifyContent={'center'} gap={5}>
                  <Text fontSize={'xl'}>{user.name}</Text>
                </Flex>
              </>
            );
          })}
        <Text fontSize={'2xl'}>bill list</Text>
        {billList &&
          billList.bills.map((bill) => {
            return (
              <>
                <Flex justifyContent={'center'} gap={5}>
                  <Text fontSize={'xl'}>
                    {bill.srcUserName} ⇒ {bill.dstUserName}
                  </Text>
                  <Text fontSize={'xl'}>￥{bill.price}</Text>
                </Flex>
              </>
            );
          })}
        <Text fontSize={'2xl'}>Sumarize</Text>
        {sumarize &&
          sumarize.bills.map((bill) => {
            return (
              <>
                <Flex justifyContent={'center'} gap={5}>
                  <Text fontSize={'xl'}>
                    {bill.srcUserName} ⇒ {bill.dstUserName}
                  </Text>
                  <Text fontSize={'xl'}>￥{bill.price}</Text>
                </Flex>
              </>
            );
          })}

        <Text fontSize={'2xl'}>new Bill</Text>
        <BillForm onRegisterComplete={fetch} selectItems={memberList ?? []} />
      </Flex>
    </>
  );
}

const getBillList = async (
  projectId: number
): Promise<BillServiceListResponse | ConnectError> => {
  try {
    const res = await Client(BillService).list({ projectId });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};

const getProjectDetail = async (
  projectId: number
): Promise<ProjectServiceGetResponse | ConnectError> => {
  try {
    const res = await Client(ProjectService).get({ id: projectId });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};

const getSumerize = async (projectId: number) => {
  try {
    const res = await Client(BillService).sumarizeByProject({ projectId });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};

const getMemberList = async (projectId: number) => {
  try {
    const res = await Client(ProjectService).joinList({ id: projectId });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};
