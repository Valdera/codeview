import {
  Box,
  createStyles,
  Group,
  Paper,
  Progress,
  SimpleGrid,
  Text,
} from '@mantine/core';
import { IconDeviceAnalytics } from '@tabler/icons';

const useStyles = createStyles((theme) => ({
  progressLabel: {
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
    lineHeight: 1,
    fontSize: theme.fontSizes.sm,
  },

  stat: {
    borderBottom: `1px solid`,
    paddingBottom: '5px',
  },

  statCount: {
    lineHeight: 1.3,
  },

  diff: {
    display: 'flex',
    alignItems: 'center',
  },

  icon: {
    color:
      theme.colorScheme === 'dark'
        ? theme.colors.dark[3]
        : theme.colors.gray[4],
  },
}));

interface ISegmentStat {
  total: string;
  data: {
    label: string;
    count: string;
    part: number;
    color: string;
  }[];
}

const SegmentStat: React.FC<ISegmentStat> = ({ total, data }) => {
  const { classes } = useStyles();

  const segments = data.map((segment) => ({
    value: segment.part,
    color: segment.color,
    label: segment.part > 10 ? `${segment.part}%` : undefined,
  }));

  const descriptions = data.map((stat) => (
    <Box
      key={stat.label}
      sx={{ borderBottomColor: stat.color }}
      className={classes.stat}
    >
      <Text
        tt={'uppercase'}
        fz={'xs'}
        c={'dimmed'}
        fw={700}
        className={'text-primary-700'}
      >
        {stat.label}
      </Text>

      <Group position={'apart'} align={'flex-end'} spacing={0}>
        <Text fw={700}>{stat.count}</Text>
        <Text c={stat.color} fw={700} size={'sm'} className={classes.statCount}>
          {stat.part}%
        </Text>
      </Group>
    </Box>
  ));

  return (
    <Paper withBorder p={'md'} radius={'md'}>
      <Group position={'apart'}>
        <Group align={'flex-end'} spacing={'xs'}>
          <Text fz={'xl'} fw={700} className={'text-primary-700'}>
            {total}
          </Text>
        </Group>
        <IconDeviceAnalytics
          size={'1.4rem'}
          className={classes.icon}
          stroke={1.5}
        />
      </Group>

      <Progress
        sections={segments}
        size={34}
        classNames={{ label: classes.progressLabel }}
        mt={20}
      />
      <SimpleGrid
        cols={3}
        breakpoints={[{ maxWidth: 'xs', cols: 1 }]}
        mt={'xl'}
      >
        {descriptions}
      </SimpleGrid>
    </Paper>
  );
};

export default SegmentStat;
