import { ChevronRightIcon } from '@chakra-ui/icons';
import { Flex, Grid, GridItem, IconButton, Text } from '@chakra-ui/react';
import TagLabel from '@components/labels/tag/TagLabel';
import { CollectionItem } from '@lib/types';
import { IconGripVertical } from '@tabler/icons';
import cn from 'classnames';
import { Emoji, EmojiStyle } from 'emoji-picker-react';
import { Draggable } from 'react-beautiful-dnd';
import s from './CollectionItemSectionCard.module.scss';

export interface ICollectionItemSectionCard {
  index: number;
  data: {
    collectionItem: CollectionItem;
  };
  disabled: boolean;
}

const MAX_NUM_TAGS = 4;

const CollectionItemSectionCard: React.FC<ICollectionItemSectionCard> = ({
  index,
  data: { collectionItem },
  disabled = false,
}) => {
  return (
    <Draggable
      isDragDisabled={disabled}
      key={collectionItem.itemId}
      index={index}
      draggableId={collectionItem.itemId}
    >
      {(provided, snapshot) => (
        <Grid
          gap={2}
          width={'full'}
          rounded={'md'}
          gridTemplateColumns={'max-content 1fr'}
          marginBottom={3}
          backgroundColor={'foreground'}
          className={cn(s.item, {
            [s.itemDragging]: snapshot.isDragging,
          })}
          ref={provided.innerRef}
          {...provided.draggableProps}
        >
          {!disabled && (
            <GridItem className={s.dragHandle} {...provided.dragHandleProps}>
              <IconGripVertical color={'white'} size={18} stroke={1.5} />
            </GridItem>
          )}
          <GridItem
            width={'full'}
            display={'grid'}
            gridTemplateColumns={'1fr max-content'}
            paddingY={'3'}
            paddingLeft={disabled ? '3' : '1'}
            paddingRight={'3'}
            colSpan={disabled ? 2 : 1}
            gap={'1'}
          >
            <Flex paddingY={'2'} width={'full'} flexDir={'column'} gap={5}>
              <Flex alignItems={'center'} gap={2}>
                <Emoji
                  size={25}
                  unified={collectionItem.emoji}
                  emojiStyle={EmojiStyle.TWITTER}
                />
                <Text fontSize={'xl'} fontWeight={'semibold'} color={'white'}>
                  {collectionItem.title}
                </Text>
              </Flex>
              <Flex width={'full'} marginTop={'auto'} flexWrap={'wrap'} gap={2}>
                {collectionItem.tags
                  .slice(0, Math.min(collectionItem.tags.length, MAX_NUM_TAGS))
                  .map((t) => (
                    <TagLabel
                      key={t.id}
                      size={'md'}
                      tag={t}
                      maxLabelLength={20}
                    />
                  ))}
              </Flex>
            </Flex>
            {disabled && (
              <IconButton
                colorScheme={'primary'}
                height={'full'}
                aria-label={'Collection Item'}
                size={'lg'}
                icon={<ChevronRightIcon fontSize={'2xl'} color={'white'} />}
              />
            )}
          </GridItem>
        </Grid>
      )}
    </Draggable>
  );
};

export default CollectionItemSectionCard;
