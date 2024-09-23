import { atom, createStore } from 'jotai';

type MeType = {
  id: number;
  name: string;
};

const store = createStore();
const MeAtom = atom<MeType | undefined>();

store.set(MeAtom, undefined);

export { MeAtom, store };
