import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Select,
  Spacer,
  useToast,
  Wrap,
  WrapItem,
} from '@chakra-ui/react';
import { useState } from 'react';
import { Client } from '../use-client';
import { BillService } from '../gen/proto_bill/v1/bill_connect';
import { ConnectError } from '@connectrpc/connect';
import { useParams } from 'react-router-dom';
import { useAtomValue } from 'jotai';
import { MeAtom } from '../utils/meAtom';

type BillFormProps = {
  onRegisterComplete: (id: string) => void;
  selectItems: { id: number; name: string }[];
};

export default function BillForm(billFromProps: BillFormProps) {
  const { id } = useParams<{ id: string }>();

  const [userId, setUserId] = useState<number | undefined>();
  const [amount, setAmount] = useState<number | undefined>();
  const [error, setError] = useState('');
  const toast = useToast();
  const me = useAtomValue(MeAtom);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!id) {
      setError('id is not provided');
      return;
    }
    if (!userId || !me?.id || (userId && userId === me?.id)) {
      console.log(userId, me?.id);
      setError('ユーザの選択が正しくありません。');
      return;
    }
    if (!amount || amount <= 0) {
      setError('金額が入力されていません');
      return;
    }

    const res = await createBill(Number(id), userId, amount, me.id);
    if (res instanceof Error) {
      setError(res.message);
      return;
    }
    billFromProps.onRegisterComplete(id);
  };

  return (
    <Wrap>
      <form onSubmit={handleSubmit}>
        <Flex direction={'column'} justify={'center'} align={'center'}>
          <FormControl>
            <FormLabel>請求先</FormLabel>
            <Select
              placeholder='Select dst'
              onChange={(e) => setUserId(Number(e.target.value))}
            >
              {billFromProps.selectItems.map((item) => {
                if (item.id !== me?.id) {
                  return <option value={item.id}>{item.name}</option>;
                }
              })}
            </Select>
          </FormControl>

          <Spacer p={5} />

          <FormControl>
            <FormLabel>金額</FormLabel>
            <Input
              type='number'
              id='amount'
              value={amount}
              onChange={(e) => setAmount(Number(e.target.value))}
              required
            />
          </FormControl>

          <Spacer p={5} />

          <WrapItem>
            <Button
              type='submit'
              onClick={() => {
                if (!amount || amount <= 0 || !userId || !me?.id) {
                  toast({
                    title: '請求作成エラー',
                    description: error,
                    status: 'error',
                    duration: 1000,
                    isClosable: true,
                  });
                } else {
                  toast({
                    title: '請求作成',
                    description: '請求を作成しました',
                    status: 'success',
                    duration: 1000,
                    isClosable: true,
                  });
                }
              }}
            >
              submit
            </Button>
          </WrapItem>
        </Flex>
      </form>
    </Wrap>
  );
}

const createBill = async (
  projectId: number,
  dstUserId: number,
  amount: number,
  srdUserId: number
) => {
  try {
    const res = await Client(BillService).create({
      projectId: projectId,
      dstUserId: dstUserId,
      price: amount,
      srcUserId: srdUserId,
    });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};
