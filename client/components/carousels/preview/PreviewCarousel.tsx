import {
  Box,
  Button,
  Flex,
  Heading,
  HStack,
  Image,
  Tag as ChakraTag,
  useBreakpointValue,
} from '@chakra-ui/react';

import { Tag } from '@lib/types/';
import { Swiper, SwiperSlide } from 'swiper/react';

import { Autoplay } from 'swiper';
import 'swiper/css';

export interface ICarouselCard {
  title: string;
  image: string;
  tags: Tag[];
}

const CarouselCard: React.FC<ICarouselCard> = ({ title, image, tags }) => {
  console.log(image);

  return (
    <Box rounded={'md'} width={'full'} height={'full'}>
      <Image
        zIndex={-2}
        width={'full'}
        height={'full'}
        position={'absolute'}
        rounded={'md'}
        src={image}
        alt={'Dan Abramov'}
      />
      <Box
        zIndex={-1}
        width={'full'}
        height={'full'}
        position={'absolute'}
        rounded={'md'}
        backgroundColor={'primary.500'}
        opacity={0.7}
      />
      <Flex flexDir={'column'} gap={3} padding={3} height={'full'}>
        <Heading color={'white'} fontSize={'3xl'}>
          {title}
        </Heading>
        <HStack>
          {tags.map((tag) => (
            <ChakraTag
              fontWeight={'bold'}
              size={'sm'}
              key={tag.id}
              backgroundColor={tag.color}
              color={'white'}
            >
              {tag.label}
            </ChakraTag>
          ))}
        </HStack>
        <Button
          size={'sm'}
          width={'100px'}
          color={'primary.500'}
          marginTop={'auto'}
        >
          See More
        </Button>
      </Flex>
    </Box>
  );
};

const data = [
  {
    image: '/assets/cover_1.png',
    title: 'Java Spring Boot',
    tags: [
      { id: '1', label: 'Array', color: '#FC7300' },
      { id: '2', label: 'Binary Tree', color: '#00425A' },
    ],
  },
  {
    image: '/assets/cover_2.png',
    title: 'Java Spring Boot',
    tags: [
      { id: '1', label: 'Array', color: '#FC7300' },
      { id: '2', label: 'Binary Tree', color: '#00425A' },
    ],
  },
  {
    image: '/assets/cover_2.png',
    title: 'Java Spring Boot',
    tags: [
      { id: '1', label: 'Array', color: '#FC7300' },
      { id: '2', label: 'Binary Tree', color: '#00425A' },
    ],
  },
  {
    image: '/assets/cover_1.png',
    title: 'Java Spring Boot',
    tags: [
      { id: '1', label: 'Array', color: '#FC7300' },
      { id: '2', label: 'Binary Tree', color: '#00425A' },
    ],
  },
  {
    image: '/assets/cover_1.png',
    title: 'Java Spring Boot',
    tags: [
      { id: '1', label: 'Array', color: '#FC7300' },
      { id: '2', label: 'Binary Tree', color: '#00425A' },
    ],
  },
  {
    image: '/assets/cover_1.png',
    title: 'Java Spring Boot',
    tags: [
      { id: '1', label: 'Array', color: '#FC7300' },
      { id: '2', label: 'Binary Tree', color: '#00425A' },
    ],
  },
];

const PreviewCarousel = () => {
  const slideWidth = useBreakpointValue(
    {
      base: '100%',
      md: '400px',
    },
    {
      ssr: true,
      fallback: '100%',
    }
  );

  const slidesPerView = useBreakpointValue(
    {
      base: 1,
      md: 2,
      lg: 1,
      xl: 2,
      '2xl': 4,
    },
    {
      ssr: true,
      fallback: '1',
    }
  );

  return (
    <>
      <Swiper
        className={'w-full h-full'}
        spaceBetween={10}
        slidesPerView={slidesPerView ?? 1}
        onSlideChange={() => console.log('slide change')}
        onSwiper={(swiper) => console.log(swiper)}
        loop={true}
        modules={[Autoplay]}
        autoplay={{
          delay: 2500,
          disableOnInteraction: true,
        }}
      >
        {data.map((item, i) => (
          <SwiperSlide key={i} style={{ width: slideWidth }}>
            <CarouselCard {...item} />
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

export default PreviewCarousel;
