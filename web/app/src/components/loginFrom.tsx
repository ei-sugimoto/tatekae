import React, { useState } from 'react';
import { Client } from '../use-client';
import { UserService } from '../gen/proto_user/v1/user_connect';
import { LoginResponse } from '../gen/proto_user/v1/user_pb';
import { ConnectError } from '@connectrpc/connect';
import {
  Button,
  FormControl,
  FormLabel,
  Input,
  Spacer,
  useToast,
  Wrap,
  WrapItem,
} from '@chakra-ui/react';

const LoginForm: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const toast = useToast();
  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    const res = await Login(email, password);

    if (res instanceof ConnectError) {
      console.error('Error:', res);
      setError(res.message);
      return;
    }

    console.log('LoginResponse:', res);
    console.log('Email:', email);
    console.log('Password:', password);
  };

  return (
    <Wrap>
      <form onSubmit={handleSubmit}>
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
              if (error) {
                toast({
                  title: 'ログインエラー',
                  description: error,
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
