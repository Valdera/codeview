import { Flex, Text } from '@chakra-ui/react';
import StatusLabel from '@components/labels/status/StatusLabel';
import TagLabel from '@components/labels/tag/TagLabel';
import { Note } from '@lib/types';
import { Emoji, EmojiStyle } from 'emoji-picker-react';
import Link from 'next/link';

const MAX_NUM_TAGS = 2;

export interface INoteListCard {
  note: Note;
}

const NoteListCard: React.FC<INoteListCard> = ({ note }) => {
  return (
    <Link href={`/note/detail/${note.id}`}>
      <Flex
        height={'full'}
        backgroundColor={'foreground'}
        padding={5}
        flexDir={'column'}
        gap={2}
        cursor={'pointer'}
        transition={'all .3s'}
        _hover={{
          transform: 'scale(1.05)',
        }}
      >
        <Flex width={'full'} alignItems={'center'}>
          <Emoji
            unified={note.emoji}
            size={22}
            emojiStyle={EmojiStyle.TWITTER}
          />
          <StatusLabel size={'sm'} status={note.status} marginLeft={'auto'} />
        </Flex>
        <Text
          noOfLines={2}
          fontSize={'xl'}
          fontWeight={'semibold'}
          color={'white'}
          marginBottom={'5'}
        >
          {note.title}
        </Text>

        <Flex
          width={'full'}
          gridTemplateColumns={{ base: 'repeat(3, 1fr)' }}
          marginTop={'auto'}
          flexWrap={'wrap'}
          gap={2}
        >
          {note.tags
            .slice(0, Math.min(note.tags.length, MAX_NUM_TAGS))
            .map((t) => (
              <TagLabel key={t.id} size={'sm'} tag={t} maxLabelLength={8} />
            ))}
        </Flex>
      </Flex>
    </Link>
  );
};

export default NoteListCard;
