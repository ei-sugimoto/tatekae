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
} from '@chakra-ui/react';

const LoginForm: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    const res = await Login(email, password);

    if (res instanceof ConnectError) {
      console.error('Error:', res);
      return;
    }

    console.log('LoginResponse:', res);
    console.log('Email:', email);
    console.log('Password:', password);
  };

  return (
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
      <Button type='submit'>Login</Button>
    </form>
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
