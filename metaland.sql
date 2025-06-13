/*
 Navicat Premium Data Transfer

 Source Server         : metaland
 Source Server Type    : MySQL
 Source Server Version : 80024
 Source Host           : 45.249.209.63:3306
 Source Schema         : metaland

 Target Server Type    : MySQL
 Target Server Version : 80024
 File Encoding         : 65001

 Date: 12/05/2025 19:35:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET GLOBAL sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- ----------------------------
-- Table structure for bounty
-- ----------------------------
DROP TABLE IF EXISTS `bounty`;
CREATE TABLE `bounty` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
  `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Transcation Hash',
  `deposit_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Contract Address',
  `startup_id` bigint NOT NULL DEFAULT '0',
  `comer_id` bigint NOT NULL DEFAULT '0',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `apply_cutoff_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `discussion_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `deposit_token_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `applicant_deposit` int NOT NULL DEFAULT '0',
  `founder_deposit` int NOT NULL DEFAULT '0',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `payment_mode` int NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `total_reward_token` int NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_applicant
-- ----------------------------
DROP TABLE IF EXISTS `bounty_applicant`;
CREATE TABLE `bounty_applicant` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `bounty_id` bigint NOT NULL DEFAULT '0',
  `comer_id` bigint NOT NULL DEFAULT '0',
  `apply_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `revoke_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `approve_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `quit_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `submit_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_contact
-- ----------------------------
DROP TABLE IF EXISTS `bounty_contact`;
CREATE TABLE `bounty_contact` (
  `id` bigint NOT NULL,
  `bounty_id` bigint NOT NULL DEFAULT '0',
  `contact_type` tinyint NOT NULL DEFAULT '0',
  `contact_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bounty_contact_uindex` (`bounty_id`,`contact_type`,`contact_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_deposit
-- ----------------------------
DROP TABLE IF EXISTS `bounty_deposit`;
CREATE TABLE `bounty_deposit` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0',
  `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `bounty_id` bigint NOT NULL DEFAULT '0',
  `comer_id` bigint NOT NULL DEFAULT '0',
  `access` int NOT NULL DEFAULT '0',
  `token_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `token_amount` int NOT NULL DEFAULT '0',
  `timestamp` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_payment_period
-- ----------------------------
DROP TABLE IF EXISTS `bounty_payment_period`;
CREATE TABLE `bounty_payment_period` (
  `id` bigint NOT NULL,
  `bounty_id` bigint NOT NULL DEFAULT '0',
  `period_type` tinyint(1) NOT NULL DEFAULT '0',
  `period_amount` bigint NOT NULL DEFAULT '0',
  `hours_per_day` int NOT NULL DEFAULT '0',
  `token1_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `token1_amount` int NOT NULL DEFAULT '0',
  `token2_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `token2_amount` int NOT NULL DEFAULT '0',
  `target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bounty_id_uindex` (`bounty_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_payment_terms
-- ----------------------------
DROP TABLE IF EXISTS `bounty_payment_terms`;
CREATE TABLE `bounty_payment_terms` (
  `id` bigint NOT NULL,
  `bounty_id` bigint NOT NULL DEFAULT '0',
  `payment_mode` tinyint(1) NOT NULL DEFAULT '0',
  `token1_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `token1_amount` int NOT NULL DEFAULT '0',
  `token2_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `token2_amount` int NOT NULL DEFAULT '0',
  `terms` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `seq_num` int NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for chain
-- ----------------------------
DROP TABLE IF EXISTS `chain`;
CREATE TABLE `chain` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain name',
  `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain logo',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-normal, 2-disable',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `chain_id` (`chain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for chain_contract
-- ----------------------------
DROP TABLE IF EXISTS `chain_contract`;
CREATE TABLE `chain_contract` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain contract address',
  `project` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1 Startup, 2 Bounty, 3 Crowdfunding, 4 Gover',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1工厂合约、2子合约',
  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'contract version',
  `abi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'abi json',
  `created_tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'created tx hash',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for chain_endpoint
-- ----------------------------
DROP TABLE IF EXISTS `chain_endpoint`;
CREATE TABLE `chain_endpoint` (
  `id` bigint NOT NULL,
  `protocol` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Communication protocol, 1-rpc 2-wss',
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain name',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-normal, 2-disable',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for comer
-- ----------------------------
DROP TABLE IF EXISTS `comer`;
CREATE TABLE `comer` (
  `id` bigint NOT NULL,
  `address` char(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'comer could save some useful info on block chain with this address',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_address_uindex` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_account
-- ----------------------------
DROP TABLE IF EXISTS `comer_account`;
CREATE TABLE `comer_account` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer unique identifier',
  `oin` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'comer outer account unique identifier, wallet will be public key and Oauth is the OauthID',
  `is_primary` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'comer use this account as primay account',
  `nick` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'comer nick name',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'avatar link address',
  `type` int NOT NULL DEFAULT '0' COMMENT '1 for github  2 for google 3 for twitter 4 for facebook 5 for likedin',
  `is_linked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 for unlink 1 for linked',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_account_oin_uindex` (`oin`),
  KEY `comer_account_comer_id_index` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_education
-- ----------------------------
DROP TABLE IF EXISTS `comer_education`;
CREATE TABLE `comer_education` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `school` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '学校名称',
  `degree` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '学位',
  `major` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '专业',
  `start_date` date DEFAULT NULL COMMENT '开始日期',
  `end_date` date DEFAULT NULL COMMENT '结束日期',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `comer_education_comer_id_index` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_skill
-- ----------------------------
DROP TABLE IF EXISTS `comer_skill`;
CREATE TABLE `comer_skill` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `skill_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '技能名称',
  `level` tinyint NOT NULL DEFAULT '0' COMMENT '熟练度(1-5)',
  `years` int NOT NULL DEFAULT '0' COMMENT '使用年限',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `comer_skill_comer_id_index` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_social
-- ----------------------------
DROP TABLE IF EXISTS `comer_social`;
CREATE TABLE `comer_social` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `platform` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '平台(twitter/discord/telegram等)',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `is_verified` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否认证',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_social_platform_username_index` (`comer_id`, `platform`, `username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_language
-- ---------------------------- 
DROP TABLE IF EXISTS `comer_language`;
CREATE TABLE `comer_language` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `language` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '语言',
  `code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '语言代码',
  `level` tinyint NOT NULL DEFAULT '0' COMMENT '熟练度(1-5)',
  `is_native` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否母语',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_language_comer_id_language_index` (`comer_id`, `language`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_follow_rel
-- ----------------------------
DROP TABLE IF EXISTS `comer_follow_rel`;
CREATE TABLE `comer_follow_rel` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comer_id` bigint NOT NULL DEFAULT '0',
  `target_comer_id` bigint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for comer_profile
-- ----------------------------
DROP TABLE IF EXISTS `comer_profile`;
CREATE TABLE `comer_profile` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'name',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'avatar',
  `cover` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `location` char(42) NOT NULL DEFAULT '' COMMENT 'location city',
  `time_zone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'time zone: UTC-09:30',
  `website` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'website',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'email',
  `twitter` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'twitter',
  `discord` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'discord',
  `telegram` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'telegram',
  `medium` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'medium',
  `facebook` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `linktree` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'bio',
  `languages` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `educations` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_profile_comer_id_uindex` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for crowdfunding
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding`;
CREATE TABLE `crowdfunding` (
  `id` bigint NOT NULL COMMENT 'crowdfunding id',
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain id',
  `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Tx hash',
  `crowdfunding_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Crowdfunding contract address',
  `startup_id` bigint NOT NULL DEFAULT '0' COMMENT 'Startup id',
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'Founder''s comer id',
  `raise_goal` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Raise goal total',
  `raise_balance` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Raise token balance',
  `sell_token_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Sell token contract address',
  `sell_token_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Sell token name',
  `sell_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Sell token symbol',
  `sell_token_decimals` int NOT NULL DEFAULT '0' COMMENT 'Sell token decimals',
  `sell_token_supply` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Sell token total supply',
  `sell_token_deposit` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Sell token deposit',
  `sell_token_balance` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Sell token balance',
  `buy_token_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Buy token contract address',
  `buy_token_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Buy token name',
  `buy_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Buy token symbol',
  `buy_token_decimals` int NOT NULL DEFAULT '0' COMMENT 'Buy token decimals',
  `buy_token_supply` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Buy token total supply',
  `team_wallet` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Team wallet address',
  `swap_percent` float NOT NULL DEFAULT '0' COMMENT 'Swap percent',
  `buy_price` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'IBO rate',
  `max_buy_amount` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Maximum buy amount',
  `max_sell_percent` float NOT NULL DEFAULT '0' COMMENT 'Maximum selling percent',
  `sell_tax` float NOT NULL DEFAULT '0' COMMENT 'Selling tax',
  `start_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'Start time',
  `end_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'End time',
  `poster` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Poster url',
  `youtube` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Youtube link',
  `detail` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Detail url',
  `description` varchar(520) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Description content',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:Pending 1:Upcoming 2:Live 3:Ended 4:Cancelled 5:Failure',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for crowdfunding_ibo_rate
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding_ibo_rate`;
CREATE TABLE `crowdfunding_ibo_rate` (
  `id` bigint NOT NULL,
  `crowdfunding_id` bigint NOT NULL DEFAULT '0' COMMENT 'Crowdfunding id',
  `end_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'End time',
  `max_buy_amount` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Maximum buy amount',
  `max_sell_percent` float NOT NULL DEFAULT '0' COMMENT 'Maximum sell percent',
  `buy_price` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'IBO rate',
  `swap_percent` float NOT NULL DEFAULT '0' COMMENT 'Swap percent',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for crowdfunding_investor
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding_investor`;
CREATE TABLE `crowdfunding_investor` (
  `id` bigint NOT NULL,
  `crowdfunding_id` bigint NOT NULL DEFAULT '0' COMMENT 'Crowdfunding id',
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'Investor'' comer id',
  `buy_token_total` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Buy token total',
  `buy_token_balance` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Buy token balance',
  `sell_token_total` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Selling token total',
  `sell_token_balance` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Selling token balance',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `crowdfunding_comer_uindex` (`crowdfunding_id`,`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for crowdfunding_swap
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding_swap`;
CREATE TABLE `crowdfunding_swap` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain id',
  `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Tx hash',
  `timestamp` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:Pending 1:Success 2:Failure',
  `crowdfunding_id` bigint NOT NULL DEFAULT '0' COMMENT 'Crowdfunding id',
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'Comer id',
  `access` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1:Invest 2:Withdraw',
  `buy_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Buy token symbol',
  `buy_token_amount` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Buy token amount',
  `sell_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Selling token symbol',
  `sell_token_amount` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Selling token amount',
  `price` decimal(38,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT 'Swap price',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for dict_data
-- ----------------------------
DROP TABLE IF EXISTS `dict_data`;
CREATE TABLE `dict_data` (
  `id` int NOT NULL,
  `startup_id` bigint NOT NULL DEFAULT '0',
  `dict_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `dict_label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `dict_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `seq_num` int NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1:enabled 2:disabled',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_admin
-- ----------------------------
DROP TABLE IF EXISTS `governance_admin`;
CREATE TABLE `governance_admin` (
  `id` int NOT NULL,
  `setting_id` bigint NOT NULL DEFAULT '0',
  `wallet_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_choice
-- ----------------------------
DROP TABLE IF EXISTS `governance_choice`;
CREATE TABLE `governance_choice` (
  `id` bigint NOT NULL,
  `proposal_id` bigint NOT NULL DEFAULT '0',
  `item_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `seq_num` tinyint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_proposal
-- ----------------------------
DROP TABLE IF EXISTS `governance_proposal`;
CREATE TABLE `governance_proposal` (
  `id` int NOT NULL,
  `startup_id` bigint NOT NULL DEFAULT '0',
  `author_comer_id` bigint NOT NULL DEFAULT '0',
  `author_wallet_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `chain_id` bigint NOT NULL DEFAULT '0',
  `block_number` bigint NOT NULL DEFAULT '0',
  `release_timestamp` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `ipfs_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `description` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `discussion_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `vote_system` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `start_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `end_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:pending 1:upcoming 2:active 3:ended',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_setting
-- ----------------------------
DROP TABLE IF EXISTS `governance_setting`;
CREATE TABLE `governance_setting` (
  `id` bigint NOT NULL,
  `startup_id` bigint NOT NULL DEFAULT '0',
  `comer_id` bigint NOT NULL DEFAULT '0',
  `vote_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `allow_member` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:no  1:yes',
  `proposal_threshold` double NOT NULL DEFAULT '0',
  `proposal_validity` double NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_strategy
-- ----------------------------
DROP TABLE IF EXISTS `governance_strategy`;
CREATE TABLE `governance_strategy` (
  `id` int NOT NULL,
  `setting_id` bigint NOT NULL DEFAULT '0',
  `dict_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `strategy_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `chain_id` bigint NOT NULL DEFAULT '0',
  `token_contract_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `vote_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `vote_decimals` int NOT NULL DEFAULT '0',
  `token_min_balance` double NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_vote
-- ----------------------------
DROP TABLE IF EXISTS `governance_vote`;
CREATE TABLE `governance_vote` (
  `id` bigint NOT NULL,
  `proposal_id` bigint NOT NULL DEFAULT '0',
  `voter_comer_id` bigint NOT NULL DEFAULT '0',
  `voter_wallet_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `choice_item_id` bigint NOT NULL DEFAULT '0',
  `choice_item_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `votes` double NOT NULL DEFAULT '0',
  `ipfs_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
  `id` bigint NOT NULL,
  `category` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'name',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'url',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `image_category_name_uindex` (`category`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for post_update
-- ----------------------------
DROP TABLE IF EXISTS `post_update`;
CREATE TABLE `post_update` (
  `id` int NOT NULL,
  `source_type` tinyint(1) NOT NULL DEFAULT '0',
  `source_id` bigint NOT NULL DEFAULT '0',
  `comer_id` bigint NOT NULL DEFAULT '0',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `timestamp` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for startup
-- ----------------------------
DROP TABLE IF EXISTS `startup`;
CREATE TABLE `startup` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer_id',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'name',
  `mode` smallint NOT NULL DEFAULT '0' COMMENT '0:NONE, 1:ESG, 2:NGO, 3:DAO, 4:COM',
  `logo` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'logo',
  `cover` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `mission` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'logo',
  `token_contract_address` char(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'token contract address',
  `overview` text NOT NULL COMMENT 'overview',
  `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `on_chain` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'whether it is on the chain',
  `kyc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'KYC',
  `contract_audit` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'contract audit',
  `website` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'website',
  `discord` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'discord',
  `twitter` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'twitter',
  `telegram` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'telegram',
  `docs` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'docs',
  `email` varchar(180) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `facebook` varchar(180) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `medium` varchar(180) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `linktree` varchar(180) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `launch_network` int NOT NULL DEFAULT '0' COMMENT 'chain id',
  `token_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'token name',
  `token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'token symbol',
  `total_supply` bigint NOT NULL DEFAULT '0' COMMENT 'total supply',
  `presale_start` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'presale start date',
  `presale_end` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'presale end date',
  `launch_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'launch_date',
  `tab_sequence` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `startup_name_uindex` (`name`),
  UNIQUE KEY `startup_token_contract_index` (`token_contract_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_follow_rel
-- ----------------------------
DROP TABLE IF EXISTS `startup_follow_rel`;
CREATE TABLE `startup_follow_rel` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer_id',
  `startup_id` bigint NOT NULL DEFAULT '0' COMMENT 'startup_id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `startup_followed_comer_id_startup_id_uindex` (`comer_id`,`startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_group
-- ----------------------------
DROP TABLE IF EXISTS `startup_group`;
CREATE TABLE `startup_group` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer_id',
  `startup_id` bigint NOT NULL DEFAULT '0' COMMENT 'startup_id',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'group name',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `startup_group_name_unidex` (`name`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_group_member_rel
-- ----------------------------
DROP TABLE IF EXISTS `startup_group_member_rel`;
CREATE TABLE `startup_group_member_rel` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer_id',
  `startup_id` bigint NOT NULL DEFAULT '0' COMMENT 'startup_id',
  `group_id` bigint NOT NULL DEFAULT '0' COMMENT 'group id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `startup_group_comer_id_uindex` (`comer_id`,`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_team_member_rel
-- ----------------------------
DROP TABLE IF EXISTS `startup_team_member_rel`;
CREATE TABLE `startup_team_member_rel` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer_id',
  `startup_id` bigint NOT NULL DEFAULT '0' COMMENT 'startup_id',
  `position` text NOT NULL COMMENT 'title',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `startup_team_rel_comer_id_startup_id_uindex` (`comer_id`,`startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_wallet
-- ----------------------------
DROP TABLE IF EXISTS `startup_wallet`;
CREATE TABLE `startup_wallet` (
  `id` bigint NOT NULL,
  `comer_id` bigint NOT NULL DEFAULT '0' COMMENT 'comer_id',
  `startup_id` bigint NOT NULL DEFAULT '0' COMMENT 'startup_id',
  `wallet_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'wallet name',
  `wallet_address` char(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'wallet address',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'name',
  `is_index` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is index',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  `category` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tag_category_name_uindex` (`name`,`category`)
) ENGINE=InnoDB AUTO_INCREMENT=114112701034523 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for tag_target_rel
-- ----------------------------
DROP TABLE IF EXISTS `tag_target_rel`;
CREATE TABLE `tag_target_rel` (
  `id` bigint NOT NULL,
  `target` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'comerSkill,startup',
  `target_id` bigint NOT NULL DEFAULT '0' COMMENT 'target id',
  `tag_id` bigint NOT NULL DEFAULT '0' COMMENT 'skill id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_id_skill_id_uindex` (`target`,`target_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0',
  `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `timestamp` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:Pending 1:Success 2:Failure',
  `source_type` tinyint(1) NOT NULL DEFAULT '0',
  `source_id` bigint NOT NULL DEFAULT '0',
  `retry_times` int NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
