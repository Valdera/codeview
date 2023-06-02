import {
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
} from '@chakra-ui/react';
import TransferList, {
  TransferListData,
} from '@components/inputs/transfer-list/TransferList';
import { useCollectionStore } from '@lib/stores';
import { CollectionItem } from '@lib/types';
import { useState } from 'react';

export interface ICollectionEditModal {
  data: {
    collectionItems: CollectionItem[];
  };
  isOpen: boolean;
  onClose: () => void;
}

const CollectionEditModal: React.FC<ICollectionEditModal> = ({
  data,
  isOpen,
  onClose,
}) => {
  const { collection, setCollectionNoteItem } = useCollectionStore();

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [itemsValue, setItemsValue] = useState<TransferListData[]>(
    collection && collection.items
      ? collection.items.map((item) => ({
          value: item.itemId,
          label: item.title,
          emoji: item.emoji,
          tags: item.tags,
        }))
      : []
  );

  const [itemsData, setItemsData] = useState<TransferListData[]>(
    data.collectionItems.map((item) => ({
      value: item.itemId,
      label: item.title,
      emoji: item.emoji,
      tags: item.tags,
    }))
  );

  return (
    <Modal
      size={{ base: 'lg', sm: 'xl', md: '2xl' }}
      onClose={onClose}
      isOpen={isOpen}
      scrollBehavior={'inside'}
      isCentered
    >
      <ModalOverlay bg={'blackAlpha.300'} backdropFilter={'blur(5px)'} />
      <ModalContent color={'white'} backgroundColor={'foreground'}>
        <ModalHeader>Edit Collection</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          {collection && (
            <TransferList
              value={itemsValue}
              data={itemsData}
              labels={['Item List', 'Current List']}
              onChange={(value) => {
                setItemsData(value[0] as TransferListData[]);
                setItemsValue(value[1] as TransferListData[]);

                console.log(value);
              }}
            />
          )}
        </ModalBody>
        <ModalFooter>
          <Button
            isLoading={isLoading}
            colorScheme={'primary'}
            onClick={() => {
              setIsLoading(true);
              setCollectionNoteItem(itemsValue.map((item) => item.value));
              setIsLoading(false);
              onClose();
            }}
          >
            Save
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default CollectionEditModal;
