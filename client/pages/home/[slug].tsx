import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import FeatureListSection from '@components/sections/feature-list/FeatureListSection';
import { NextPageWithLayout } from '@pages/page';

const HomePage: NextPageWithLayout = () => {
  return (
    <>
      <FeatureListSection />
    </>
  );
};

HomePage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export default HomePage;
