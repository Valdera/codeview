import { Note, NoteItem, NoteMetadata } from '@lib/types';
import { StateCreator } from 'zustand';

export interface NoteSlice {
  note: Note | null;
  load: (note: Note) => void;
  updateNote: (body: NoteMetadata) => void;
  createNoteItem: () => void;
  updateNoteItem: (
    id: string,
    body: { header: string; content: string }
  ) => void;
  deleteNoteItem: (id: string) => void;
  reorderNoteItem: (from: number, to: number) => void;
}

export const createNoteSlice: StateCreator<NoteSlice> = (set, get) => {
  return {
    note: null,
    load: (note) => {
      set({ note });
    },
    updateNote: ({ title, tags }) => {
      const note = get().note;
      if (!note) return;

      // TODO: Update Note through backend

      note.title = title;
      note.tags = tags;

      set({ note });
    },
    createNoteItem: () => {
      const note = get().note;
      if (!note) return;
      if (!note.items) note.items = [];

      // TODO: Create note item through backend

      const noteItem: NoteItem = {
        id: Date.now().toString(),
        header: 'Untitled',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Heading</strong></h2>`,
        numOrder: note.items.length,
      };

      note.items.push(noteItem);

      set({ note });
    },
    updateNoteItem: (id, { header, content }) => {
      const note = get().note;
      if (!note || !note.items) return;

      // TODO: Update note item through backend

      note.items = note.items.map((item) => {
        if (item.id == id) {
          return { ...item, header: header, content: content };
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

      console.log(
        'BEFORE: ',
        note.items.forEach((t) => console.log(t))
      );

      const items = [...note.items];
      const [removed] = items.splice(from, 1);
      items.splice(to, 0, removed);
      note.items = items;

      console.log(
        'AFTER: ',
        note.items.forEach((t) => console.log(t))
      );

      set({ note });
    },
  };
};
