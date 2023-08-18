CREATE TABLE IF NOT EXISTS t_activity (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    activity_name varchar(100) NOT NULL comment '活动名称',
    start_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '开始时间',
    end_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '结束时间',
    introduce text DEFAULT NULL comment '活动介绍',
    requirements text DEFAULT NULL comment '活动需求'
    );

CREATE TABLE IF NOT EXISTS t_nft_airdrop(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    fk_activity_id int NOT NULL  comment '活动id',
    wallet_address varchar(100) NOT NULL comment '钱包地址',
    deploy_network varchar(64) DEFAULT NULL comment '部署网络',
    contract_name varchar(64) DEFAULT NULL comment '合约名字',
    contract_address varchar(64) DEFAULT NULL comment '合约地址',
    contract_abi text DEFAULT NULL comment '合约abi',
    deploy_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '部署时间'
    );

INSERT INTO media.t_activity (id, activity_name, start_time, end_time, introduce, requirements) VALUES (1, 'Hamster&Scroll联名NFT空投', '2023-08-22 00:00:00', '2023-09-23 00:00:00', '活动介绍', '在活动时间内，在Hamster将合约部署到Scroll主网。');

