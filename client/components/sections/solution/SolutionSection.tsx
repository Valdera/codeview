import { DeleteIcon } from '@chakra-ui/icons';
import { Box, Heading, Highlight, HStack, IconButton } from '@chakra-ui/react';
import RichTextEditor from '@components/rte/RichTextEditor';
import { useProblemStore } from '@lib/stores';
import { Solution } from '@lib/types/problem';

export interface ISolutionSection {
  solution: Solution;
}

const SolutionSection: React.FC<ISolutionSection> = ({ solution }) => {
  const { deleteSolution, updateSolution } = useProblemStore();

  return (
    <Box width={'full'} padding={'5'} backgroundColor={'white'}>
      <HStack
        width={'full'}
        marginBottom={5}
        display={'flex'}
        alignItems={'center'}
      >
        <Heading as={'h3'} size={'md'}>
          <Highlight
            query={'solution'}
            styles={{ px: '2', py: '1', rounded: 'full', bg: 'green.100' }}
          >
            Solution
          </Highlight>
        </Heading>
        <IconButton
          className={'!ml-auto'}
          size={'sm'}
          aria-label={'delete solution'}
          colorScheme={'red'}
          icon={<DeleteIcon />}
          onClick={() => deleteSolution(solution.id)}
        />
      </HStack>
      <RichTextEditor
        content={solution.content}
        onSave={(content) => {
          updateSolution({ id: solution.id, content: content });
        }}
      />
    </Box>
  );
};

export default SolutionSection;
