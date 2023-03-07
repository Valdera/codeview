import { createNoteSlice, NoteSlice } from '@lib/slices/noteSlice';
import { create } from 'zustand';

type StoreState = NoteSlice;

const useNoteStore = create<StoreState>()((...a) => ({
  ...createNoteSlice(...a),
}));

export default useNoteStore;
