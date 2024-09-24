import React, { useState } from 'react';
import { Client } from '../use-client';
import { UserService } from '../gen/proto_user/v1/user_connect';
import { LoginResponse } from '../gen/proto_user/v1/user_pb';
import { ConnectError } from '@connectrpc/connect';
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
import { useNavigate } from 'react-router-dom';
import { useAtom } from 'jotai';
import { MeAtom } from '../utils/meAtom';

export const LoginForm: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const toast = useToast();
  const navigate = useNavigate();

  const [me, setMe] = useAtom(MeAtom);

  const handleSubmit = async (event: React.FormEvent) => {
    setError('');
    event.preventDefault();
    const res = await Login(email, password);

    if (res instanceof ConnectError) {
      console.error('Error:', res);
      setError(res.message);
      return;
    }
    localStorage.setItem('token', res.token);
    setMe((prev) => ({ ...prev, id: res.id, name: res.username }));
    console.log(me);
    console.log(res);
    navigate('/dashboard');

    return;
  };

  return (
    <Wrap>
      <form onSubmit={handleSubmit}>
        <Flex direction={'column'} justify={'center'} align={'center'}>
          <div>
            <FormControl>
              <FormLabel>Email:</FormLabel>
              <Input
                type='email'
                id='email'
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </FormControl>
          </div>
          <div>
            <FormControl>
              <FormLabel>Password:</FormLabel>
              <Input
                type='password'
                id='password'
                value={password}
                onChange={(e) => setPassword(e.target.value)}
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
                    title: 'ログインエラー',
                    description: 'パスワードかメールアドレスが間違っています',
                    status: 'error',
                    duration: 1000,
                    isClosable: true,
                  });
                }
              }}
            >
              Login
            </Button>
          </WrapItem>
        </Flex>
      </form>
    </Wrap>
  );
};

export default LoginForm;

const Login = async (
  email: string,
  password: string
): Promise<LoginResponse | ConnectError> => {
  try {
    const res = await Client(UserService).login({
      email: email,
      password: password,
    });
    return res;
  } catch (e) {
    const connectErr = ConnectError.from(e);

    return connectErr;
  }
};
