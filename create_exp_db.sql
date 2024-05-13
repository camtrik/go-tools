CREATE DATABASE IF NOT EXISTS test_examples;
USE test_examples;

CREATE TABLE IF NOT EXISTS `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT 'Tag name',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'Creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modification time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Deletion time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'Deleted status, 0 for not deleted, 1 for deleted',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'Status, 0 for disabled, 1 for enabled',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Tag management';

CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
