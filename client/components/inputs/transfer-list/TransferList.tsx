import { Flex, Text } from '@chakra-ui/react';
import TagLabel from '@components/labels/tag/TagLabel';
import { Tag } from '@lib/types';
import {
  Checkbox,
  Group,
  TransferList as TransferListMantine,
  TransferListData as TransferListDataMantine,
  TransferListItemComponent,
  TransferListItemComponentProps,
} from '@mantine/core';
import { Emoji, EmojiStyle } from 'emoji-picker-react';

const MAX_NUM_TAGS = 2;

export interface TransferListData {
  value: string;
  label: string;
  emoji: string;
  tags: Tag[];
}

export interface ITransferList {
  value?: TransferListData[];
  data?: TransferListData[];
  labels?: [string, string];
  onChange?: (value: TransferListDataMantine) => void;
}

const TransferList: React.FC<ITransferList> = ({
  value = [],
  data = [],
  labels = ['', ''],
  onChange = () => {},
}) => {
  return (
    <TransferListMantine
      value={[data, value]}
      onChange={onChange}
      searchPlaceholder={'Search...'}
      nothingFound={'Nothing found'}
      titles={labels}
      listHeight={300}
      breakpoint={'sm'}
      itemComponent={ItemComponent}
      filter={(query, item) =>
        item.label.toLowerCase().includes(query.toLowerCase().trim()) ||
        item.description.toLowerCase().includes(query.toLowerCase().trim())
      }
      styles={(theme) => ({
        transferListBody: {
          backgroundColor: theme.colors.background,
          borderColor: theme.colors.foreground,
        },
        transferListHeader: {
          backgroundColor: theme.colors.background,
        },
        transferListSearch: {
          backgroundColor: theme.colors.background,
          borderColor: theme.colors.foreground,

          '&:focus': {
            borderColor: theme.colors.primary[5],
          },
        },
        transferListControl: {
          borderColor: theme.colors.foreground,
          color: theme.colors.gray[2],

          '&:disabled': {
            borderColor: theme.colors.foreground,
            color: theme.colors.gray[7],
          },

          '&:hover': {
            backgroundColor: theme.colors.gray[7],
          },
        },
      })}
    />
  );
};

const ItemComponent: TransferListItemComponent = ({
  data,
  selected,
}: TransferListItemComponentProps) => (
  <Group noWrap>
    <Emoji size={30} unified={data.emoji} emojiStyle={EmojiStyle.TWITTER} />
    <div style={{ flex: 1 }}>
      <Text color={'gray.200'} fontSize={'md'} marginBottom={'1'}>
        {data.label}
      </Text>

      <Flex width={'full'} marginTop={'auto'} flexWrap={'wrap'} gap={2}>
        {data.tags
          .slice(0, Math.min(data.tags.length, MAX_NUM_TAGS))
          .map((t: Tag) => (
            <TagLabel key={t.id} size={'sm'} tag={t} maxLabelLength={8} />
          ))}
      </Flex>
    </div>
    <Checkbox
      checked={selected}
      onChange={() => {}}
      tabIndex={-1}
      sx={{ pointerEvents: 'none' }}
    />
  </Group>
);

export default TransferList;
