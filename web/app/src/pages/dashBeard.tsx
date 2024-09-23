import { ConnectError } from '@connectrpc/connect';
import { ProjectService } from '../gen/proto_project/v1/project_connect';
import {
  ProjectServiceJoinResponse,
  ProjectServiceListResponse,
} from '../gen/proto_project/v1/project_pb';
import { Client } from '../use-client';
import { useEffect, useState } from 'react';
import { Button, Flex, Heading, Spacer, Text } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export function DashBoard() {
  const [list, setList] = useState<ProjectServiceListResponse>();
  const [error, setError] = useState<string | null>(null);

  const fetchList = async () => {
    const res = await getProjectList();
    if (res instanceof ConnectError) {
      setError(res.message);
      return;
    }
    setList(res);
  };

  useEffect(() => {
    fetchList();
  }, []);

  if (error) {
    return (
      <>
        <h1>Error: {error}</h1>
      </>
    );
  }

  return (
    <>
      <Heading mb={5}>Project List</Heading>
      <Flex
        justifyContent={'center'}
        alignItems={'center'}
        direction={'column'}
      >
        {list?.projects.map((project) => {
          return (
            <Flex
              w={'max-content'}
              justifyContent={'center'}
              alignContent={'center'}
            >
              <>
                <Link to={`/project/${project.id}`}>
                  <Text fontSize={'xl'}>{project.name}</Text>
                </Link>
                <Spacer p={8} />
              </>
              <>
                <Button
                  onClick={() => {
                    joinProject(project.id);
                  }}
                >
                  Join
                </Button>
              </>
            </Flex>
          );
        })}
      </Flex>
    </>
  );
}

const getProjectList = async (): Promise<
  ProjectServiceListResponse | ConnectError
> => {
  try {
    const res = await Client(ProjectService).list({});

    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};

const joinProject = async (
  projectId: number
): Promise<ProjectServiceJoinResponse | ConnectError> => {
  try {
    const res = await Client(ProjectService).join({ id: projectId });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};
