import { Note, NoteItem } from '@lib/types';
import { StateCreator } from 'zustand';

export interface NoteSlice {
  note: Note | null;
  load: (note: Note) => void;
  updateNote: (body: {
    title: string;
    tags: string[];
    emoji: string;
    status: string;
  }) => void;
  createNoteItem: () => void;
  updateNoteItem: (id: string, body: { content: string }) => void;
  deleteNoteItem: (id: string) => void;
  reorderNoteItem: (from: number, to: number) => void;
}

export const createNoteSlice: StateCreator<NoteSlice> = (set, get) => {
  return {
    note: null,
    load: (note) => {
      set({ note });
    },
    updateNote: (value) => {
      const note = get().note;
      if (!note) return;

      // TODO: Update Note through backend

      // note.title = title;
      // note.tags = tags;
      // note.emoji = emoji;
      // note.status = status;

      console.log('value to be post: ', value);

      set({ note });
    },
    createNoteItem: () => {
      const note = get().note;
      if (!note) return;
      if (!note.items) note.items = [];

      // TODO: Create note item through backend

      const noteItem: NoteItem = {
        id: Date.now().toString(),
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Heading</strong></h2>`,
        numOrder: note.items.length,
      };

      note.items.push(noteItem);

      set({ note });
    },
    updateNoteItem: (id, { content }) => {
      const note = get().note;
      if (!note || !note.items) return;

      // TODO: Update note item through backend

      note.items = note.items.map((item) => {
        if (item.id == id) {
          return { ...item, content: content };
        }
        return item;
      });

      set({ note });
    },
    deleteNoteItem: (id) => {
      const note = get().note;
      if (!note || !note.items) return;

      // TODO: Delete note item through backend

      note.items = note.items.filter((item) => item.id != id);

      set({ note });
    },
    reorderNoteItem: (from, to) => {
      const note = get().note;
      if (!note || !note.items) return;

      // TODO: Reorder note item through backend

      const items = [...note.items];
      const [removed] = items.splice(from, 1);
      items.splice(to, 0, removed);
      note.items = items;

      set({ note });
    },
  };
};
