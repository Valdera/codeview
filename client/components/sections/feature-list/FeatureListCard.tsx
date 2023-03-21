import {
  Box,
  Button,
  Flex,
  Heading,
  Stack,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import Link from 'next/link';
import { ReactElement } from 'react';

interface IFeatureListCard {
  heading: string;
  description: string;
  icon: ReactElement;
  href: string;
}

export const FeatureListCard: React.FC<IFeatureListCard> = ({
  heading,
  description,
  icon,
  href,
}) => {
  return (
    <Box
      maxW={{ base: 'full', md: '275px' }}
      w={'full'}
      borderWidth={'1px'}
      borderRadius={'lg'}
      backgroundColor={'white'}
      overflow={'hidden'}
      p={5}
    >
      <Stack align={'start'} spacing={2}>
        <Flex
          w={16}
          h={16}
          align={'center'}
          justify={'center'}
          color={'white'}
          rounded={'full'}
          bg={useColorModeValue('gray.100', 'gray.700')}
        >
          {icon}
        </Flex>
        <Box mt={2}>
          <Heading size={'md'} color={'primary.700'}>
            {heading}
          </Heading>
          <Text mt={1} fontSize={'sm'}>
            {description}
          </Text>
        </Box>
        <Link href={href}>
          <Button variant={'link'} colorScheme={'secondary'} size={'sm'}>
            Learn more
          </Button>
        </Link>
      </Stack>
    </Box>
  );
};

export default FeatureListCard;
