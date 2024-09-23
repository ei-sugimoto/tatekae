import { Flex, Heading, Spacer, Text } from '@chakra-ui/react';
import { useParams } from 'react-router-dom';
import {
  BillServiceListResponse,
  BillServiceSumrizeByProjectResponse,
} from '../gen/proto_bill/v1/bill_pb';
import { ConnectError } from '@connectrpc/connect';
import { Client } from '../use-client';
import { BillService } from '../gen/proto_bill/v1/bill_connect';
import { useEffect, useState } from 'react';
import { ProjectServiceGetResponse } from '../gen/proto_project/v1/project_pb';
import { ProjectService } from '../gen/proto_project/v1/project_connect';

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
    useState<BillServiceSumrizeByProjectResponse | null>(null);

  useEffect(() => {
    if (!id) {
      setError(new ConnectError('id is not provided'));
      return;
    }
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

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
  };

  return (
    <>
      <Heading>{project?.name}</Heading>
      <Spacer mt={5} />
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
