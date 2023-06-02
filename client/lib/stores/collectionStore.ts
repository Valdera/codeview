import { create } from 'zustand';
import {
  CollectionSlice,
  createCollectionSlice,
} from '../slices/collectionSlice';

type StoreState = CollectionSlice;

const useCollectionStore = create<StoreState>()((...a) => ({
  ...createCollectionSlice(...a),
}));

export default useCollectionStore;
