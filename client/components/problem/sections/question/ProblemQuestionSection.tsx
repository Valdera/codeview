import { Box, Heading, Highlight, HStack } from '@chakra-ui/react';
import RichTextEditor from '@components/rte/RichTextEditor';
import { useProblemStore } from '@lib/stores';
import { Question } from '@lib/types';

export interface IProblemQuestionSection {
  question: Question;
  disabled?: boolean;
}

const ProblemQuestionSection: React.FC<IProblemQuestionSection> = ({
  question,
  disabled = false,
}) => {
  const { updateQuestion } = useProblemStore();

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
            query={'question'}
            styles={{ px: '2', py: '1', rounded: 'full', bg: 'teal.100' }}
          >
            Question
          </Highlight>
        </Heading>
      </HStack>
      <RichTextEditor
        content={question.content}
        onSave={(content) => {
          updateQuestion({
            content,
          });
        }}
        disabled={disabled}
      />
    </Box>
  );
};

export default ProblemQuestionSection;
