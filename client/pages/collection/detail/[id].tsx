import { Flex, VStack } from '@chakra-ui/react';
import CollectionItemSection from '@components/collection/sections/item/CollectionItemSection';
import CollectionMetadataSection from '@components/collection/sections/metadata/CollectionMetadataSection';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import { useCollectionStore } from '@lib/stores';
import { Collection } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { useEffect } from 'react';

interface ICollectionDetailPage {
  collection: Collection;
}

const CollectionDetailPage: NextPageWithLayout<ICollectionDetailPage> = ({
  collection: initialCollection,
}) => {
  const { load, collection } = useCollectionStore();

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
        <VStack gap={5} width={{ base: '100%', md: '95%', lg: '80%' }}>
          <CollectionMetadataSection data={{}} disabled={true} />
          <CollectionItemSection disabled={true} />
        </VStack>
      </Flex>
    </>
  );
};

CollectionDetailPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<ICollectionDetailPage>> => {
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

  return {
    props: {
      collection: MOCK_COLLECTION,
    },
  };
};

export default CollectionDetailPage;
