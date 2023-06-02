package config

type GCloudConfig struct {
	BucketName string `yaml:"bucketName" envconfig:"GC_BUCKET_NAME" default:""`
}
