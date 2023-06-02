import {
  Box,
  Button,
  Flex,
  Heading,
  HStack,
  Tag as ChakraTag,
} from '@chakra-ui/react';

import { Tag } from '@lib/types/';

export interface ICollectionCarouselCard {
  title: string;
  tags: Tag[];
}

const CollectionCarouselCard: React.FC<ICollectionCarouselCard> = ({
  title,
  tags,
}) => {
  return (
    <Box rounded={'md'} width={'full'} height={'full'}>
      <Box
        zIndex={-1}
        width={'full'}
        height={'full'}
        position={'absolute'}
        rounded={'md'}
        backgroundColor={'primary.600'}
      />
      <Flex flexDir={'column'} gap={3} padding={3} height={'full'}>
        <Heading color={'white'} fontSize={'3xl'}>
          {title}
        </Heading>
        <HStack>
          {tags.map((tag) => (
            <ChakraTag
              fontWeight={'bold'}
              size={'sm'}
              key={tag.id}
              backgroundColor={tag.color}
              color={'white'}
            >
              {tag.label}
            </ChakraTag>
          ))}
        </HStack>
        <Button
          size={'sm'}
          width={'100px'}
          color={'primary.500'}
          marginTop={'auto'}
        >
          See More
        </Button>
      </Flex>
    </Box>
  );
};

export default CollectionCarouselCard;
