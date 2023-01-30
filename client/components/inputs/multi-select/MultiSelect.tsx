import { Flex, Text } from '@chakra-ui/react';
import {
  CloseButton,
  MultiSelect as MantineSelect,
  MultiSelectValueProps,
} from '@mantine/core';

/**TODO:
 * - Add Custom value
 */

export interface IMultiSelect {
  data: { value: string; label: string; color: string }[];
  value?: string[];
  creatable?: boolean;
  label?: string;
  disabled?: boolean;
  handleChange?: (value: string[]) => void;
  handleCreate?: ((query: string) => string | null) | undefined;
}

const MultiSelect: React.FC<IMultiSelect> = ({
  data,
  handleCreate,
  handleChange,
  value = [],
  disabled = false,
  creatable = true,
  label = '',
}) => {
  return (
    <>
      <MantineSelect
        label={label}
        data={data}
        defaultValue={value}
        placeholder={'Select items'}
        searchable
        creatable={creatable}
        valueComponent={(props) => {
          return <Value color={props.color} {...props} />;
        }}
        readOnly={disabled}
        transitionDuration={150}
        transition={'pop-top-left'}
        transitionTimingFunction={'ease'}
        getCreateLabel={(query) => `+ Create ${query}`}
        onCreate={handleCreate}
        onChange={handleChange}
        className={'w-full'}
        styles={(_themes) => ({
          values: {},
          input: {
            margin: '2px',
            transform: 'scale(1)',
          },
          wrapper: {
            transform: 'scale(1)',
          },
          searchInput: {
            margin: 0,
          },
        })}
      />
    </>
  );
};

const Value: React.FC<
  MultiSelectValueProps & { value: string; color: string }
> = ({ color, label, onRemove, ...others }) => {
  return (
    <div className={others.className}>
      <Flex
        alignItems={'center'}
        backgroundColor={color}
        padding={'2px'}
        paddingLeft={'3'}
        borderRadius={'md'}
      >
        <Text
          color={'white'}
          lineHeight={'1'}
          fontSize={'12'}
          fontWeight={'medium'}
        >
          {label}
        </Text>
        <CloseButton
          onMouseDown={onRemove}
          variant={'transparent'}
          size={20}
          iconSize={12}
          tabIndex={-1}
          className={'text-white font-medium'}
        />
      </Flex>
    </div>
  );
};

export default MultiSelect;
