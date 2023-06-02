import { Collection } from '@lib/types';
import { StateCreator } from 'zustand';

export interface CollectionSlice {
  collection: Collection | null;
  load: (collection: Collection) => void;
  updateCollection: (body: {
    title: string;
    description: string;
    tags: string[];
    emoji: string;
  }) => void;
  setCollectionNoteItem: (itemIds: string[]) => void;
  reorderCollectionItem: (from: number, to: number) => void;
}

export const createCollectionSlice: StateCreator<CollectionSlice> = (
  set,
  get
) => {
  return {
    collection: null,
    load: (collection) => {
      set({ collection });
    },
    updateCollection: (value) => {
      const collection = get().collection;
      if (!collection) return;

      console.log('value to be post', value);
      set({ collection });
    },
    setCollectionNoteItem: (itemIds) => {
      const collection = get().collection;
      if (!collection || !collection.items) return;

      console.log(itemIds);
    },
    reorderCollectionItem: (from, to) => {
      const collection = get().collection;
      if (!collection || !collection.items) return;

      // TODO: Reorder note item through backend

      const items = [...collection.items];
      const [removed] = items.splice(from, 1);
      items.splice(to, 0, removed);
      collection.items = items;

      set({ collection });
    },
  };
};
