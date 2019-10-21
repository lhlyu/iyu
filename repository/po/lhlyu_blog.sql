/*
SQLyog Ultimate v12.08 (64 bit)
MySQL - 5.6.28-cdb2016-log : Database - lhlyu_blog
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`lhlyu_blog` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin */;

USE `lhlyu_blog`;

/*Table structure for table `yu_article` */

DROP TABLE IF EXISTS `yu_article`;

CREATE TABLE `yu_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '赞',
  `unlike` int(11) NOT NULL DEFAULT '0' COMMENT '踩',
  `view` int(11) NOT NULL DEFAULT '0' COMMENT '浏览量',
  `comments_number` int(11) NOT NULL DEFAULT '0' COMMENT '评论数量',
  `bg` varchar(200) NOT NULL DEFAULT '' COMMENT '头背景',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶:0-不置顶;1-置顶',
  `nail_id` int(11) NOT NULL DEFAULT '0' COMMENT '钉子ID',
  `kind` tinyint(1) NOT NULL DEFAULT '0' COMMENT '文章类型:0-普通文章;1-特殊文章',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已删除:0-未删除;1-已删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章';

/*Table structure for table `yu_article_catalog` */

DROP TABLE IF EXISTS `yu_article_catalog`;

CREATE TABLE `yu_article_catalog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL DEFAULT '0',
  `catalog_id` int(11) NOT NULL,
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0-未删除；1-已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='文章分类关联表';

/*Table structure for table `yu_article_label` */

DROP TABLE IF EXISTS `yu_article_label`;

CREATE TABLE `yu_article_label` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL DEFAULT '0',
  `label_id` int(11) NOT NULL,
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0-未删除；1-已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='文章标签关联表';

/*Table structure for table `yu_category` */

DROP TABLE IF EXISTS `yu_category`;

CREATE TABLE `yu_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT '分类名字',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除:0-未删除;1-已删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类';

/*Table structure for table `yu_comment` */

DROP TABLE IF EXISTS `yu_comment`;

CREATE TABLE `yu_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `article_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `floor` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '楼层',
  `content` text COLLATE utf8_bin NOT NULL COMMENT '评论内容',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '赞',
  `unlike` int(11) NOT NULL DEFAULT '0' COMMENT '踩',
  `is_check` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论是否已审核:0-未审核;1-已审核',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论是否已被删除:0-未删除;1-已删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='评论表';

/*Table structure for table `yu_label` */

DROP TABLE IF EXISTS `yu_label`;

CREATE TABLE `yu_label` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT '标签名字',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0-未删除；1-已删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签';

/*Table structure for table `yu_nail` */

DROP TABLE IF EXISTS `yu_nail`;

CREATE TABLE `yu_nail` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '钉子ID',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT '钉子名字',
  `color` varchar(7) NOT NULL DEFAULT '' COMMENT '钉子颜色',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除: 0-未删除；1-已删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='钉子';

/*Table structure for table `yu_post` */

DROP TABLE IF EXISTS `yu_post`;

CREATE TABLE `yu_post` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '评论ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `floor` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '楼层',
  `at_id` int(11) NOT NULL DEFAULT '0' COMMENT '艾特回复的ID',
  `at_floor` varchar(20) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '艾特回复的楼层',
  `content` text COLLATE utf8_bin NOT NULL COMMENT '评论内容',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '赞',
  `unlike` int(11) NOT NULL DEFAULT '0' COMMENT '踩',
  `is_check` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论是否已审核:0-未审核;1-已审核',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论是否已被删除:0-未删除;1-已删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='评论回复表';

/*Table structure for table `yu_quanta` */

DROP TABLE IF EXISTS `yu_quanta`;

CREATE TABLE `yu_quanta` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `key` varchar(20) NOT NULL DEFAULT '' COMMENT 'key值',
  `value` varchar(200) NOT NULL DEFAULT '' COMMENT 'value值',
  `desc` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `is_enable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启用:0-启用;1-不启用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='配置表';

/*Table structure for table `yu_user` */

DROP TABLE IF EXISTS `yu_user`;

CREATE TABLE `yu_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `third_id` int(11) NOT NULL DEFAULT '0' COMMENT '第三方登录返回的ID',
  `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是管理员:0-普通;1-观察者;9-管理员',
  `from` int(2) NOT NULL DEFAULT '0' COMMENT '来源:1-github;2-gitee',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户状态:0-正常;1-已删除;2-黑名单',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Table structure for table `yu_user_info` */

DROP TABLE IF EXISTS `yu_user_info`;

CREATE TABLE `yu_user_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `avatar_url` varchar(200) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户头像',
  `user_url` varchar(200) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户地址',
  `user_name` varchar(50) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名字',
  `bio` varchar(200) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '个性签名',
  `ip` varchar(50) COLLATE utf8_bin NOT NULL DEFAULT '0.0.0.0' COMMENT 'ip',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='用户详情表';

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
