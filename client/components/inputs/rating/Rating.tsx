import { Flex, Text } from '@chakra-ui/react';
import { Rating as MantineRating } from '@mantine/core';

export interface IRating {
  value?: number;
  label?: string;
  handleChange?: (value: number) => void;
}

const Rating: React.FC<IRating> = ({
  value = 0,
  handleChange = () => {},
  label = '',
}) => {
  return (
    <div>
      {label && (
        <Text fontSize={'sm'} fontWeight={'500'} fontFamily={'heading'}>
          Rating
        </Text>
      )}
      <Flex height={10} alignItems={'center'}>
        <MantineRating value={value} onChange={handleChange} />
      </Flex>
    </div>
  );
};

export default Rating;
