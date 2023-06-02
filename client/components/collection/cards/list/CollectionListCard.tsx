import { Button, Flex, Text } from '@chakra-ui/react';
import TagLabel from '@components/labels/tag/TagLabel';
import { Collection } from '@lib/types';
import { Emoji, EmojiStyle } from 'emoji-picker-react';
import Link from 'next/link';

const MAX_NUM_TAGS = 4;

export interface ICollectionListCard {
  collection: Collection;
}

const CollectionListCard: React.FC<ICollectionListCard> = ({ collection }) => {
  return (
    <Flex
      backgroundColor={'foreground'}
      padding={5}
      flexDir={'column'}
      gap={5}
      borderRadius={'md'}
    >
      <Flex width={'full'} alignItems={'center'} gap={2}>
        <Emoji
          unified={collection.emoji}
          size={22}
          emojiStyle={EmojiStyle.TWITTER}
        />
        <Text
          noOfLines={2}
          fontSize={'xl'}
          fontWeight={'semibold'}
          color={'white'}
        >
          {collection.title}
        </Text>
      </Flex>

      <Flex width={'full'} marginTop={'auto'} flexWrap={'wrap'} gap={2}>
        {collection.tags
          .slice(0, Math.min(collection.tags.length, MAX_NUM_TAGS))
          .map((t) => (
            <TagLabel key={t.id} size={'md'} tag={t} maxLabelLength={20} />
          ))}
      </Flex>

      <Text color={'gray.100'} noOfLines={3} width={'full'}>
        {collection.description}
      </Text>
      <Link href={`/collection/detail/${collection.id}`} className={'ml-auto'}>
        <Button colorScheme={'primary'} size={'md'}>
          See More
        </Button>
      </Link>
    </Flex>
  );
};

export default CollectionListCard;
