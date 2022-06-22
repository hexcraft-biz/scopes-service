-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/

--
-- 資料庫: `oauth2-scopes`
--

-- --------------------------------------------------------

--
-- 傾印資料表的資料 `scopes`
--

INSERT INTO `scopes` (`id`, `resource_domain_name`, `resource_name`, `name`, `type`, `ctime`, `mtime`) VALUES
(0x0b41b735b737475597392d48515eca57, 'http://account-backend.default.svc.cluster.local/', 'user.socialmedia', 'user.socialmedia.readonly', 'public', '2022-06-22 13:16:05', '2022-06-22 13:23:26'),
(0x0f5412e8428d4ee9aa29c888710266a1, 'http://account-backend.default.svc.cluster.local/', 'user.contact', 'user.contact.readonly', 'public', '2022-06-22 13:13:50', '2022-06-22 13:23:26'),
(0x13708e0d24aa4be4a24aef7d820c7182, 'http://account-backend.default.svc.cluster.local/', 'user.preference', 'user.preference', 'public', '2022-06-22 13:14:13', '2022-06-22 13:23:26'),
(0x1692cc73c3674f7481c35976c4fdc2f9, 'http://account-backend.default.svc.cluster.local/', 'users.account', 'user.account.readonly', 'public', '2022-06-01 08:49:43', '2022-06-22 13:23:26'),
(0x215fde721cae45acb894e684f7892afa, 'http://account-backend.default.svc.cluster.local/', 'user.socialmedia', 'user.socialmedia', 'public', '2022-06-22 13:15:56', '2022-06-22 13:23:26'),
(0x23ad3a49e9464cb3a42b750c4cf80894, 'http://account-backend.default.svc.cluster.local/', 'user.preference', 'user.preference.self', 'public', '2022-06-22 13:14:37', '2022-06-22 13:23:26'),
(0x27b5ad4c6b3949058e94d4a7b272d2f1, 'http://account-backend.default.svc.cluster.local/', 'users.account', 'user.account.self', 'public', '2022-06-01 08:49:47', '2022-06-22 13:23:26'),
(0x2a5fac859ae6471391fad896ce6433dd, 'http://account-backend.default.svc.cluster.local/', 'user.wallet', 'user.wallet', 'public', '2022-06-22 13:14:50', '2022-06-22 13:23:26'),
(0x2b772d6bdbbb46e08519b5ef2e4df959, 'http://account-backend.default.svc.cluster.local/', 'user.preference', 'user.preference.readonly', 'public', '2022-06-22 13:14:29', '2022-06-22 13:23:26'),
(0x62cc1eb44122455cb9d2df7f2ab418f6, 'http://account-backend.default.svc.cluster.local/', 'user.wallet', 'user.wallet.readonly', 'public', '2022-06-22 13:15:01', '2022-06-22 13:23:26'),
(0x6e4d1c4564ed4a468a63fe4f749d9429, 'http://account-backend.default.svc.cluster.local/', 'user.contact', 'user.contact.self', 'public', '2022-06-22 13:14:00', '2022-06-22 13:23:26'),
(0x8054706efa16447c91e8e1869c4086fd, 'http://account-backend.default.svc.cluster.local/', 'user.socialmedia', 'user.socialmedia.self', 'public', '2022-06-22 13:16:08', '2022-06-22 13:23:26'),
(0x93917f4a677b40269616a380a66e2f84, 'http://account-backend.default.svc.cluster.local/', 'user.contact', 'user.contact', 'public', '2022-06-22 13:13:38', '2022-06-22 13:23:26'),
(0x99dfe2670d514136b0154e7ce1558ce2, 'http://account-backend.default.svc.cluster.local/', 'users.prototype', 'user.prototype.self', 'private', '2022-06-01 08:49:14', '2022-06-22 13:23:26'),
(0x9b329b69964e47be8fb4b0bd1a74041d, 'http://account-backend.default.svc.cluster.local/', 'users.account', 'user.account', 'public', '2022-06-01 08:49:28', '2022-06-22 13:23:26'),
(0x9bf8466cb8ba4453a041b49bb9564982, 'http://account-backend.default.svc.cluster.local/', 'user.wallet', 'user.wallet.self', 'public', '2022-06-22 13:15:12', '2022-06-22 13:23:26'),
(0xa39271fb6dfa49829af7a4a1c7a73f11, 'http://account-backend.default.svc.cluster.local/', 'users.prototype', 'user.prototype', 'private', '2022-06-01 08:49:07', '2022-06-22 13:23:26');

--
-- 已傾印資料表的索引
--
