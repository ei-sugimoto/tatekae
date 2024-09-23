import React, { useState } from 'react';
import { Client } from '../use-client';
import { UserService } from '../gen/proto_user/v1/user_connect';
import { ConnectError } from '@connectrpc/connect';
import { RegisterResponse } from '../gen/proto_user/v1/user_pb';
import {
  Button,
  FormControl,
  FormLabel,
  Input,
  Spacer,
} from '@chakra-ui/react';

const RegisterForm: React.FC = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Handle form submission logic here
    const res = Register(username, email, password);
    if (res instanceof ConnectError) {
      console.error('Error:', res);
      return;
    }
    console.log('RegisterResponse:', res);
    console.log({ username, email, password });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <FormControl>
          <FormLabel>Username:</FormLabel>
          <Input
            type='text'
            id='username'
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </FormControl>
      </div>
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
      <Button type='submit'>Register</Button>
    </form>
  );
};

export default RegisterForm;

const Register = async (
  username: string,
  email: string,
  password: string
): Promise<RegisterResponse | ConnectError> => {
  try {
    const res = await Client(UserService).register({
      username: username,
      email: email,
      password: password,
    });
    return res;
  } catch (e) {
    const connectErr = ConnectError.from(e);
    return connectErr;
  }
};
