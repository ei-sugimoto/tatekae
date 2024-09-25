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
  useToast,
  Wrap,
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { useSetAtom } from 'jotai';
import { MeAtom } from '../utils/meAtom';
import { setCookie } from '../utils/cookie';

const RegisterForm: React.FC = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const setMe = useSetAtom(MeAtom);
  const toast = useToast();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const res = await Register(username, email, password);
    if (res instanceof ConnectError) {
      toast({
        title: '登録エラー',
        description: 'すでに登録済みかサーバにエラーが発生しました。',
        status: 'error',
        duration: 1000,
        isClosable: true,
      });

      return;
    }

    setCookie('token', res.token, 1);
    setMe({ id: res.id, name: res.username });
    toast({
      title: '登録成功',
      description: '登録に成功しました',
      status: 'success',
      duration: 1000,
      isClosable: true,
    });
    navigate('/dashboard');
    return;
  };

  return (
    <Wrap>
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
    </Wrap>
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
