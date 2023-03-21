import { Tag } from '@chakra-ui/react';
import { Tag as TagType } from '@lib/types';
import { truncateString } from '@lib/utils';

export interface ILabelTag {
  tag: TagType;
  defaultBg?: string | undefined;
}

const LabelTag: React.FC<ILabelTag> = ({ tag, defaultBg = undefined }) => {
  return (
    <Tag
      key={tag.label}
      width={'100px'}
      display={'flex'}
      backgroundColor={defaultBg ? defaultBg : tag.color}
      color={'white'}
      textOverflow={'ellipsis'}
      justifyContent={'center'}
    >
      {truncateString(tag.label, 8)}
    </Tag>
  );
};

export default LabelTag;
