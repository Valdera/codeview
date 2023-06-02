import { useBreakpointValue } from '@chakra-ui/react';
import { Autoplay } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';

import { ReactNode } from 'react';
import 'swiper/css';

const DEFAULT_SLIDES_WIDTH: Partial<Record<string, string>> = {
  base: '100%',
  md: '400px',
};

const DEFAULT_SLIDES_PER_VIEW: Partial<Record<string, number>> = {
  base: 1,
  md: 2,
  lg: 1,
  xl: 2,
  '2xl': 4,
};

export interface IPreviewCarousel {
  slidesPerView?: Partial<Record<string, number>>;
  slidesWidth?: Partial<Record<string, string>>;
  contents: ReactNode[];
  onSlideChange?: () => void;
  onSwiper?: () => void;
}

const PreviewCarousel: React.FC<IPreviewCarousel> = ({
  contents,
  slidesPerView = DEFAULT_SLIDES_PER_VIEW,
  slidesWidth = DEFAULT_SLIDES_WIDTH,
  onSlideChange = () => {},
  onSwiper = () => {},
}) => {
  const slidesWidthValue = useBreakpointValue(slidesWidth, {
    ssr: true,
    fallback: '100%',
  });

  const slidesPerViewValue = useBreakpointValue(slidesPerView, {
    ssr: true,
    fallback: '1',
  });

  return (
    <>
      <Swiper
        className={'w-full h-full'}
        spaceBetween={10}
        slidesPerView={slidesPerViewValue ?? 1}
        onSlideChange={onSlideChange}
        onSwiper={onSwiper}
        loop={true}
        modules={[Autoplay]}
        autoplay={{
          delay: 3000,
          disableOnInteraction: true,
        }}
      >
        {contents.map((content, i) => (
          <SwiperSlide key={i} style={{ width: slidesWidthValue }}>
            {content}
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

export default PreviewCarousel;
