import { Grid } from '@chakra-ui/react';
import { ReactElement } from 'react';
import HomeFeatureListCard from './HomeFeatureListCard';

export interface IHomeFeatureListSection {
  features: {
    heading: string;
    description: string;
    icon: ReactElement;
    href: string;
  }[];
}

const HomeFeatureListSection: React.FC<IHomeFeatureListSection> = ({
  features,
}) => {
  return (
    <Grid
      width={'full'}
      gridTemplateColumns={{
        base: 'repeat(1, 1fr)',
        sm: 'repeat(2, 1fr)',
        md: 'repeat(2, 1fr)',
        lg: 'repeat(4, 1fr)',
        xl: 'repeat(5, 1fr)',
      }}
      gap={3}
    >
      {features.map((f) => (
        <HomeFeatureListCard
          key={f.heading}
          heading={f.heading}
          icon={f.icon}
          description={f.description}
          href={f.href}
        />
      ))}
    </Grid>
  );
};

export default HomeFeatureListSection;
