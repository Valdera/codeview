import { DeleteIcon } from '@chakra-ui/icons';
import { Box, Heading, Highlight, IconButton } from '@chakra-ui/react';
import RichTextEditor from '@components/inputs/rte/RichTextEditor';
import { useProblemStore } from '@lib/stores';
import { Solution } from '@lib/types/problem';

export interface IProblemSolutionSection {
  solution: Solution;
  disabled?: boolean;
}

const ProblemSolutionSection: React.FC<IProblemSolutionSection> = ({
  solution,
  disabled = false,
}) => {
  const { deleteSolution, updateSolution } = useProblemStore();

  return (
    <Box width={'full'} padding={'5'} backgroundColor={'foreground'}>
      <RichTextEditor
        content={solution.content}
        onSave={(content) => {
          updateSolution(solution.id, { content: content });
        }}
        extraTools={
          !disabled
            ? [
                <IconButton
                  key={'delete'}
                  size={'md'}
                  aria-label={'delete solution'}
                  colorScheme={'red'}
                  icon={<DeleteIcon />}
                  onClick={() => deleteSolution(solution.id)}
                />,
              ]
            : []
        }
        heading={
          <Heading as={'h3'} size={'md'}>
            <Highlight
              query={'solution'}
              styles={{ px: '2', py: '1', rounded: 'full', bg: 'green.100' }}
            >
              Solution
            </Highlight>
          </Heading>
        }
        disabled={disabled}
      />
    </Box>
  );
};

export default ProblemSolutionSection;
