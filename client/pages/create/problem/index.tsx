import { AddIcon } from '@chakra-ui/icons';
import { Button, VStack } from '@chakra-ui/react';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import MetadataSection from '@components/sections/metadata/MetadataSection';
import QuestionSection from '@components/sections/question/QuestionSection';
import SolutionSection from '@components/sections/solution/SolutionSection';
import { useProblemStore } from '@lib/stores';
import { NextPageWithLayout } from '@pages/page';

const CreateProblemPage: NextPageWithLayout = () => {
  const { problem, createSolution } = useProblemStore();

  return (
    <>
      {problem && (
        <VStack width={'full'} gap={5}>
          <MetadataSection />
          <QuestionSection />
          {problem.solutions.map((solution) => (
            <SolutionSection key={solution.id} solution={solution} />
          ))}
          <Button
            onClick={() => createSolution()}
            rightIcon={<AddIcon />}
            colorScheme={'purple'}
            width={'full'}
          >
            Add Solution
          </Button>
        </VStack>
      )}
    </>
  );
};

export default CreateProblemPage;

CreateProblemPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};
