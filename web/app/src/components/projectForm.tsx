import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Spacer,
  useToast,
  Wrap,
  WrapItem,
} from '@chakra-ui/react';
import React, { useState } from 'react';
import { Client } from '../use-client';
import { ProjectService } from '../gen/proto_project/v1/project_connect';

type ProjectFormProps = {
  onRegisterComplete: () => void;
};

export const ProjectForm: React.FC<ProjectFormProps> = ({
  onRegisterComplete,
}) => {
  const [name, setName] = useState('');

  const [error, setError] = useState('');
  const toast = useToast();
  const handleSubmit = async (e: React.FormEvent) => {
    setError('');
    e.preventDefault();
    if (name === '') {
      setError('プロジェクト名かプロジェクト説明が入力されていません');
      return;
    }
    const res = await createProject(name);
    if (res instanceof Error) {
      setError(res.message);
      return;
    }
    setName('');
    onRegisterComplete();
  };

  return (
    <Wrap>
      <form onSubmit={handleSubmit}>
        <Flex direction={'column'} justify={'center'} align={'center'}>
          <div>
            <FormControl>
              <FormLabel>Project Name:</FormLabel>
              <Input
                type='text'
                id='name'
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
              />
            </FormControl>
          </div>

          <Spacer p={5} />
          <WrapItem>
            <Button
              type='submit'
              onClick={() => {
                if (error != '') {
                  toast({
                    title: 'プロジェクト作成エラー',
                    description:
                      'プロジェクト名かプロジェクト説明が入力されていません',
                    status: 'error',
                    duration: 1000,
                    isClosable: true,
                  });
                } else {
                  toast({
                    title: 'プロジェクト作成',
                    description: 'プロジェクトを作成しました',
                    status: 'success',
                    duration: 1000,
                    isClosable: true,
                  });
                }
              }}
            >
              Create
            </Button>
          </WrapItem>
        </Flex>
      </form>
    </Wrap>
  );
};

const createProject = async (name: string) => {
  const res = await Client(ProjectService).create({ name: name });
  return res;
};
