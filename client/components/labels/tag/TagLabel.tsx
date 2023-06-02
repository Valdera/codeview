import { Tag, TagProps } from '@chakra-ui/react';
import { Tag as TagType } from '@lib/types';
import { truncateString } from '@lib/utils';

export interface ITagLabel extends TagProps {
  tag: TagType;
  bgColor?: string | undefined;
  maxLabelLength?: number;
}

const TagLabel: React.FC<ITagLabel> = ({
  tag,
  bgColor = undefined,
  maxLabelLength = 8,
  ...props
}) => {
  return (
    <Tag
      key={tag.label}
      display={'flex'}
      backgroundColor={bgColor ? bgColor : tag.color}
      color={'white'}
      textOverflow={'ellipsis'}
      justifyContent={'center'}
      {...props}
    >
      {truncateString(tag.label, maxLabelLength)}
    </Tag>
  );
};

export default TagLabel;
