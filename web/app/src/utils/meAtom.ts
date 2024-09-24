import { atom } from 'jotai';

type MeType = {
  id: number;
  name: string;
};

const MeAtom = atom<MeType | undefined>();

export { MeAtom };
