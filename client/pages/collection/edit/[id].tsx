import { EditIcon } from '@chakra-ui/icons';
import { Flex, IconButton, useDisclosure, VStack } from '@chakra-ui/react';
import CollectionEditModal from '@components/collection/modals/edit/CollectionEditModal';
import CollectionItemSection from '@components/collection/sections/item/CollectionItemSection';
import CollectionMetadataSection from '@components/collection/sections/metadata/CollectionMetadataSection';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import { useCollectionStore } from '@lib/stores';
import { Collection, CollectionItem, Tag } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { useEffect } from 'react';

interface ICollectionEditPage {
  collectionItems: CollectionItem[];
  collection: Collection;
  tags: Tag[];
}

const CollectionEditPage: NextPageWithLayout<ICollectionEditPage> = ({
  collection: initialCollection,
  collectionItems,
  tags,
}) => {
  const { load, collection } = useCollectionStore();
  const { isOpen, onOpen, onClose } = useDisclosure();

  useEffect(() => {
    load(initialCollection);
  }, [load, initialCollection]);

  if (!collection) return <></>;

  return (
    <>
      <Flex
        width={'full'}
        alignItems={'center'}
        flexDir={'column'}
        marginBottom={{ base: '20', md: '0' }}
      >
        <VStack gap={5} width={{ base: '100%', md: '80%' }}>
          <CollectionMetadataSection data={{ tags: tags }} />
          <CollectionItemSection />
        </VStack>
      </Flex>

      <CollectionEditModal
        data={{ collectionItems: collectionItems }}
        isOpen={isOpen}
        onClose={onClose}
      />

      <IconButton
        aria-label={'Edit collection'}
        size={'lg'}
        colorScheme={'primary'}
        rounded={'100%'}
        right={{ base: '50%', md: '10' }}
        bottom={'5'}
        zIndex={'5'}
        transform={{ base: 'translateX(50%)', md: '' }}
        position={'fixed'}
        icon={<EditIcon fontSize={'xl'} />}
        onClick={onOpen}
      />
    </>
  );
};

CollectionEditPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<ICollectionEditPage>> => {
  const MOCK_COLLECTION: Collection = {
    id: '1',
    title: 'Spring Data JPA with Hibernate',
    emoji: '1f4d1',
    type: 'NOTE',
    description:
      'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae dolores, possimus pariatur animi temporibus nesciunt praesentium dolore sed nulla ipsum eveniet corporis quidem, mollitia itaque minus soluta, voluptates neque explicabo tempora nisi culpa eiusatque dignissimos. Molestias explicabo corporis voluptatem?',
    tags: [
      { id: '1', label: 'Spring Boot', color: '#FC7300' },
      { id: '2', label: 'Go', color: '#00425A' },
      { id: '3', label: 'Java', color: '#BFDB38' },
      { id: '4', label: 'Micronaut', color: '#F55050' },
      {
        id: '5',
        label: 'Clean Architecture',
        color: '#FF78F0',
      },
      {
        id: '6',
        label: 'Clojure',
        color: '#243763',
      },
    ],
    items: [
      {
        itemId: '2',
        title: 'Dependency Injection',
        emoji: '1f4d1',
        tags: [
          { id: '1', label: 'Spring Boot', color: '#FC7300' },
          { id: '3', label: 'Java', color: '#BFDB38' },
          {
            id: '5',
            label: 'Clean Architecture',
            color: '#FF78F0',
          },
          {
            id: '6',
            label: 'Clojure',
            color: '#243763',
          },
        ],
      },
      {
        itemId: '3',
        title: 'Resilient Architecture',
        emoji: '1f4d1',
        tags: [
          {
            id: '5',
            label: 'Clean Architecture',
            color: '#FF78F0',
          },
          {
            id: '6',
            label: 'Clojure',
            color: '#243763',
          },
        ],
      },
      {
        itemId: '5',
        title: 'Java Persistence',
        emoji: '1f4d1',
        tags: [
          {
            id: '6',
            label: 'Clojure',
            color: '#243763',
          },
        ],
      },
    ],
  };

  const MOCK_COLLECTION_ITEMS: CollectionItem[] = [
    {
      itemId: '2314',
      title: 'Hello World',
      emoji: '1f4d1',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
        {
          id: '5',
          label: 'Clean Architecture',
          color: '#FF78F0',
        },
        {
          id: '6',
          label: 'Clojure',
          color: '#243763',
        },
      ],
    },
    {
      itemId: '3421',
      title: 'Basic Java',
      emoji: '1f4d1',
      tags: [
        {
          id: '5',
          label: 'Clean Architecture',
          color: '#FF78F0',
        },
        {
          id: '6',
          label: 'Clojure',
          color: '#243763',
        },
      ],
    },
  ];

  const MOCK_TAGS: Tag[] = [
    { id: '1', label: 'Spring Boot', color: '#FC7300' },
    { id: '2', label: 'Go', color: '#00425A' },
    { id: '3', label: 'Java', color: '#BFDB38' },
    { id: '4', label: 'Micronaut', color: '#F55050' },
    {
      id: '5',
      label: 'Clean Architecture',
      color: '#FF78F0',
    },
    {
      id: '6',
      label: 'Clojure',
      color: '#243763',
    },
  ];

  return {
    props: {
      collection: MOCK_COLLECTION,
      tags: MOCK_TAGS,
      collectionItems: MOCK_COLLECTION_ITEMS,
    },
  };
};

export default CollectionEditPage;
