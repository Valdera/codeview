import { Tag, TagProps } from '@chakra-ui/react';

export type Status = 'DRAFT' | 'PUBLISHED';

export interface IStatusLabel extends TagProps {
  status: Status;
  colorScheme?: string | undefined;
}

const StatusLabel: React.FC<IStatusLabel> = ({
  status,
  colorScheme = undefined,
  ...props
}) => {
  const statusColorMap: { [key: string]: string } = {
    DRAFT: 'yellow',
    PUBLISHED: 'teal',
  };

  return (
    <Tag
      key={status}
      display={'flex'}
      colorScheme={colorScheme ? colorScheme : statusColorMap[status]}
      justifyContent={'center'}
      {...props}
    >
      {status}
    </Tag>
  );
};

export default StatusLabel;
