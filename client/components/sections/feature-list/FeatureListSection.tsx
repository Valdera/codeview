import {
  Box,
  Container,
  Flex,
  Heading,
  Icon,
  Stack,
  Text,
} from '@chakra-ui/react';
import {
  FcAbout,
  FcAssistant,
  FcCollaboration,
  FcDonate,
  FcManager,
} from 'react-icons/fc';
import FeatureListCard from './FeatureListCard';

const FeatureListSection = () => {
  return (
    <Box width={'full'}>
      <Stack spacing={4} as={Container} maxW={'3xl'} textAlign={'center'}>
        <Heading
          color={'primary.500'}
          fontSize={{ base: '2xl', sm: '4xl' }}
          fontWeight={'bold'}
        >
          List of Features
        </Heading>
        <Text color={'gray.600'} fontSize={{ base: 'sm', sm: 'lg' }}>
          See all of my shared resources (notes, problem solutions, blogs, and
          many more) related with software engineering.
        </Text>
      </Stack>

      <Container maxW={'5xl'} mt={12}>
        <Flex flexWrap={'wrap'} gridGap={6} justify={'center'}>
          <FeatureListCard
            heading={'Problems'}
            icon={<Icon as={FcAssistant} w={10} h={10} />}
            description={'Solutions of programming interview questions'}
            href={'/problem/list'}
          />
          <FeatureListCard
            heading={'Notes'}
            icon={<Icon as={FcCollaboration} w={10} h={10} />}
            description={'Notes about software engineering (BE, FE, ML, etc)'}
            href={'/note/list'}
          />
          <FeatureListCard
            heading={'Tutorials'}
            icon={<Icon as={FcDonate} w={10} h={10} />}
            description={"Valdera's tutorial videos on Youtube"}
            href={'/tutorial/list'}
          />
          <FeatureListCard
            heading={'Blogs'}
            icon={<Icon as={FcManager} w={10} h={10} />}
            description={"Valdera's published blogs on Medium"}
            href={'/blog/list'}
          />
          <FeatureListCard
            heading={'Collections'}
            icon={<Icon as={FcAbout} w={10} h={10} />}
            description={'Collections of multiple source'}
            href={'/collection/list'}
          />
        </Flex>
      </Container>
    </Box>
  );
};

export default FeatureListSection;
