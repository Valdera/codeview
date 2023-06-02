import { Box, Button, Flex, Grid, GridItem, VStack } from '@chakra-ui/react';
import SearchBar from '@components/inputs/search/SearchBar';
import Select from '@components/inputs/select/Select';
import { Formik } from 'formik';

interface FormValue {
  title: string;
  type: string;
}

export interface ICollectionFilterSection {
  onSubmit?: (value: FormValue) => void;
}

const CollectionFilterSection: React.FC<ICollectionFilterSection> = ({
  onSubmit = () => {},
}) => {
  const initialValues: FormValue = {
    title: '',
    type: 'ALL',
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, _action) => {
        onSubmit(values);
      }}
    >
      {(props) => (
        <form className={'w-full'} onSubmit={props.handleSubmit}>
          <Flex
            width={'full'}
            alignItems={'center'}
            justifyContent={'center'}
            paddingY={'10'}
            backgroundColor={'primary.600'}
          >
            <VStack
              width={'full'}
              alignItems={'center'}
              justifyContent={'center'}
              gridTemplateRows={'1fr 1fr'}
              gap={5}
            >
              <Box width={{ base: '80%', md: '70%', lg: '50%' }}>
                <SearchBar />
              </Box>
              <Grid
                gap={5}
                width={{ base: '70%', lg: '40%' }}
                gridTemplateColumns={'1fr max-content'}
                alignItems={'center'}
                justifyContent={'center'}
              >
                <GridItem width={'full'}>
                  <Select
                    data={[
                      {
                        value: 'NOTE',
                        label: 'Note',
                        color: 'yellow',
                      },
                      {
                        value: 'PROBLEM',
                        label: 'Problem',
                        color: 'red',
                      },
                      {
                        value: 'ALL',
                        label: 'All',
                        color: 'teal',
                      },
                    ]}
                    value={props.values.type}
                    label={''}
                    onChange={(value) => props.setFieldValue('type', value)}
                    placeholder={'Select status'}
                  />
                </GridItem>
                <Button
                  backgroundColor={'foreground'}
                  color={'primary.400'}
                  transition={'all .15s'}
                  _hover={{
                    backgroundColor: 'background',
                  }}
                >
                  Search
                </Button>
              </Grid>
            </VStack>
          </Flex>
        </form>
      )}
    </Formik>
  );
};

export default CollectionFilterSection;
