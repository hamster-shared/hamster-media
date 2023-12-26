CREATE TABLE IF NOT EXISTS t_navbar (
    id              int           NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`          varchar(100)  NOT NULL                           comment 'name',
    icon            varchar(200)      NULL                           comment 'icon',
    code            int           NOT NULL                           comment 'code',
    parent_code     int           NOT NULL                           comment 'parent_code',
    `level`         int           NOT NULL                           comment 'level',
    `path`         varchar(200)       NULL                           comment 'path',
    create_time     timestamp         NULL DEFAULT CURRENT_TIMESTAMP comment 'create_time'
);

CREATE TABLE IF NOT EXISTS t_navbar_content (
    id              int           NOT NULL AUTO_INCREMENT PRIMARY KEY,
    navbar_id       int           NOT NULL                           comment 'navbar_id',
    icon            varchar(200)      NULL                           comment 'icon',
    title           varchar(100)  NOT NULL                           comment 'title',
    content         varchar(255)  NOT NULL                           comment 'content',
    version         varchar(100)  NOT NULL                           comment 'version',
    new_flag        tinyint(1)    NOT NULL DEFAULT 0                 comment 'new flag',
    `path`          varchar(150)  NOT NULL                           comment 'font_end path',
    create_time     timestamp         NULL DEFAULT CURRENT_TIMESTAMP comment 'create_time'
);

ALTER TABLE t_navbar add sort         int    NOT NULL DEFAULT 0                 COMMENT 'sort' AFTER level;

INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (1, 'Features', null, 1, 0, 1, 1, null, '2023-11-03 17:49:01');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (2, 'About', null, 2, 0, 1, 5, null, '2023-11-03 17:49:14');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (3, 'Docs', null, 3, 0, 1, 6, null, '2023-11-03 17:49:29');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (4, 'Community', null, 4, 0, 1, 4, null, '2023-11-03 17:49:41');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (5, 'Smart Contract', 'https://g.alpha.hamsternet.io/ipfs/QmbuyPRcusBJQVZbuiUX2JttVtUy2avccmB7iBPN54EJS3', 5, 1, 2, 2, null, '2023-11-03 17:50:36');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (6, 'Front End', 'https://g.alpha.hamsternet.io/ipfs/QmRX8YNApkrtwbS5csARQXwGyZQwEcJ7NKMJT5CWMdWmcR', 6, 1, 2, 3, null, '2023-11-03 17:51:02');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (7, 'Node', 'https://g.alpha.hamsternet.io/ipfs/QmSmDRuatkxV4M6mXfSBN17mqcwG6dNNhteccSoQsjVE3p', 7, 1, 2, 4, null, '2023-11-03 17:51:23');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (8, 'Market', 'https://g.alpha.hamsternet.io/ipfs/QmdWX5L9aVMFQBwn1jgTAMXpr4pGrDQ5TiY3QxmhYvXVYb', 8, 1, 2, 1, null, '2023-11-03 17:51:40');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (9, 'Development', null, 9, 5, 3, 1, null, '2023-11-03 17:52:11');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (10, 'Secure Code', null, 10, 5, 3, 2, null, '2023-11-03 17:52:35');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (11, 'Secure Deploy', null, 11, 5, 3, 3, null, '2023-11-03 17:52:57');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (12, 'Development', null, 12, 6, 3, 1, null, '2023-11-03 17:53:35');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (13, 'Secure Code', null, 13, 6, 3, 2, null, '2023-11-03 17:53:52');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (14, 'Secure Deploy', null, 14, 6, 3, 3, null, '2023-11-03 17:54:17');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (15, 'Development', null, 15, 7, 3, 1, null, '2023-11-03 17:54:36');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (16, 'Secure Deploy', null, 16, 7, 3, 2, null, '2023-11-03 17:54:55');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (17, 'Template', null, 17, 8, 3, 1, null, '2023-11-03 17:55:30');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (18, 'Middleware', null, 18, 8, 3, 2, null, '2023-11-03 17:55:53');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (19, 'Faucet', null, 19, 0, 1, 2, null, '2023-11-03 17:49:01');
INSERT INTO media.t_navbar (id, name, icon, code, parent_code, level, sort, path, create_time) VALUES (20, 'Incubator', null, 20, 0, 1, 3, null, '2023-11-03 17:49:01');


INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (1, 9, 'https://g.alpha.hamsternet.io/ipfs/QmX5dHsSGRTtNzyaSHsHswaRjsChiniVaXMykGVa6XcJBx', 'Build Services', 'Intelligent contract building simplified. Construct smart contracts with one click, and efficiently manage versions and track code', '0.14.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Build EVM Contract', '2023-11-06 10:25:23');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (2, 9, 'https://g.alpha.hamsternet.io/ipfs/QmPy8K1tWeTyPnETisiTjQUuwMWDbSvT8YDSCucyHkEpc8', 'Sandbox Services', 'The sandbox environment can be started online for developers and the OP network that can fork different environments', 'Coming Soon', 0, 'NA', '2023-11-06 10:25:54');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (3, 9, 'https://g.alpha.hamsternet.io/ipfs/Qma3z9GKJ8xGUX41rFieFdd47YEcAzRQYTFbEkqqUPDxzg', 'ALine', 'Hamster automates your Web3.0 project workflows seamlessly. Check, build, and deploy your code directly through code repository', '0.15.0', 1, 'https://hamsternet.io/workflow', '2023-11-06 10:41:18');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (4, 10, 'https://g.alpha.hamsternet.io/ipfs/QmSXwt5ZP7TZrGkFLR4ejeabCH9pJQ5Lf6tNjAqVm8ZGXm', 'Security Analysis', 'Ensure contract security with mainstream vulnerability detection methods including static analysis of vulnerability rules and formal verification', '0.7.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Check EVM Contract', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (5, 10, 'https://g.alpha.hamsternet.io/ipfs/QmYoCNp2WWrpevjtK3jRR3CVBn7JJkxs7Fwz4V4VaL1UgX', 'Code Quality Analysis', 'Identify coding issues, potential bugs and non-compliant code in smart contract through static analysis of the source code', '0.7.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Check EVM Contract', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (6, 10, 'https://g.alpha.hamsternet.io/ipfs/QmQjYEb2S9zwdw11tjfiFbQmwxXB5xknW6uAQRCC94YLY5', 'Open Source Analysis', 'Effectively identify, monitor and manage software components and dependencies to better manage risks and ensure application security', '0.7.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Check EVM Contract', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (7, 10, 'https://g.alpha.hamsternet.io/ipfs/QmQA1czjiSD4DrmjBwjvPTXNwyJa7PHtFAfHP2GzjQ7sp2', 'Gas Usage Analysis', 'Collect and analyze gas usage, method invocations and metrics for each unit test to provide data support for optimizing gas utilization efficiency', '0.7.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Check EVM Contract', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (8, 10, 'https://g.alpha.hamsternet.io/ipfs/QmUmtezhucCSa6S8ctj8Qyfzs6aXUZPTuoanB2sjkC6Aij', 'AI Analysis', 'Leverage AI to identify potential issues and defects in smart contracts, improving their quality and security', '0.7.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Check EVM Contract', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (9, 11, 'https://g.alpha.hamsternet.io/ipfs/Qmf4j7KzY6V6ooKee4539d4hcs4ptxastmCMqoi8AgZMNy', 'Deploy Services', 'One-click deploy a single smart contract, no need to set up a private key, safe and reliable', '0.7.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Deploy EVM Contract', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (10, 11, 'https://g.alpha.hamsternet.io/ipfs/Qma2WuVobmBPNXabCCnzrvKQAbDKjBvSgLWv9Dt3KsLe86', 'Deploy Services Pro', 'One-click deployment of multiple smart contracts, supports deployment orchestration and configuration management', 'Coming Soon', 0, 'NA', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (11, 12, 'https://g.alpha.hamsternet.io/ipfs/QmX5dHsSGRTtNzyaSHsHswaRjsChiniVaXMykGVa6XcJBx', 'Build Services', 'Simplify front-end project building. One-click build of front-end code, efficiently manage versions, track code', '0.12.0', 0, 'https://hamsternet.io/docs/Automated workflow/Front-End Project/Build Front-End Code', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (12, 12, 'https://g.alpha.hamsternet.io/ipfs/Qma3z9GKJ8xGUX41rFieFdd47YEcAzRQYTFbEkqqUPDxzg', 'ALine', 'Hamster automates your Web3.0 project workflows seamlessly. Check, build, and deploy your code directly through code repository', '0.15.0', 1, 'https://hamsternet.io/workflow', '2023-11-06 10:42:21');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (13, 13, 'https://g.alpha.hamsternet.io/ipfs/QmYoCNp2WWrpevjtK3jRR3CVBn7JJkxs7Fwz4V4VaL1UgX', 'Code Quality Analysis', 'Through static analysis of source code, identify coding issues, potential errors and non-compliant code in front-end code', '0.2.0', 0, 'https://hamsternet.io/docs/Automated workflow/Front-End Project/Check Front-End Code', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (14, 14, 'https://g.alpha.hamsternet.io/ipfs/QmYQacR3MFcf171fAA2qu9CeVgJqBaYkPv4HvM632TxMyo', 'IPFS Deploy', 'Instantly deploy your front-end code to IPFS with one click', '0.2.0', 0, 'https://hamsternet.io/docs/Automated workflow/Front-End Project/Deploy Front-End Code#deploy-front-end-on-ipfs', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (15, 14, 'https://g.alpha.hamsternet.io/ipfs/QmVotVfmv7NJXQaJsbTthRqEBzYpv4KqmjrwBZRa7UXDPo', 'Container Deploy', 'Instantly deploy your front-end code to container with one click', '0.4.0', 0, 'https://hamsternet.io/docs/Automated workflow/Front-End Project/Deploy Front-End Code#deploy-front-end-on-container', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (16, 14, 'https://g.alpha.hamsternet.io/ipfs/QmVDB76zMsovQtmY4uRREkAyWN4gobugf9YZo7ieTF6GPb', 'IC Deploy', 'Deploy your front-end code to containers on the Internet Computer with one click', '0.12.0', 0, 'https://hamsternet.io/docs/Automated workflow/Front-End Project/Deploy Front-End Code#deploy-front-end-on-icinternet-computer', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (17, 15, 'https://g.alpha.hamsternet.io/ipfs/QmX5dHsSGRTtNzyaSHsHswaRjsChiniVaXMykGVa6XcJBx', 'Build Services', 'Simplify front-end project building. One-click build of Node code, efficiently manage versions, track code', '0.12.0', 0, 'NA', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (18, 15, 'https://g.alpha.hamsternet.io/ipfs/Qma3z9GKJ8xGUX41rFieFdd47YEcAzRQYTFbEkqqUPDxzg', 'ALine', 'Hamster automates your Web3.0 project workflows seamlessly. Check, build, and deploy your code directly through code repository', '0.15.0', 1, 'https://hamsternet.io/workflow', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (19, 16, 'https://g.alpha.hamsternet.io/ipfs/Qmf4j7KzY6V6ooKee4539d4hcs4ptxastmCMqoi8AgZMNy', 'Deploy Services', 'One-click deployment of chain nodes, no need to set up private keys, safe and reliable', '0.15.0', 0, 'NA', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (20, 17, 'https://g.alpha.hamsternet.io/ipfs/QmVG3ea8PkarN25pgsE6ci9z9LUtLrz9LKyQkGfUewUhue', 'Contract Template', 'The contract template provides the basic framework of a smart contract, allowing developers to quickly build customized smart contracts based on it', '0.15.0', 0, 'https://hamsternet.io/docs/Automated workflow/EVM Contract/Create Project for EVM#create-evm-project-with-template', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (21, 17, 'https://g.alpha.hamsternet.io/ipfs/QmYDb6JVnS1bc5xZvVnWpdoMAfZWm1Yf56fDcmuN1ukaBb', 'FrontEnd Template', 'Frontend code templates provide common page structures, styles and interaction logic to help frontend developers quickly scaffold the framework of a website or application', '0.2.0', 0, 'https://hamsternet.io/docs/Automated workflow/Front-End Project/Create Project for Front-End#create-front-end-project-with-template', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (22, 17, 'https://g.alpha.hamsternet.io/ipfs/QmZW3R6FftxoviiDc7i9Q9fVh5L6E1wFzdHConz59LooQf', 'Node Template', 'Node templates provide standard project structure, dependency configuration, etc. to help developers quickly initialize a node project', '0.12.0', 0, 'NA', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (23, 17, 'https://g.alpha.hamsternet.io/ipfs/QmQNRSMTUFxTJ6WsR8kXLn4ZaJMR12mciK2oKiqfi1TS6Z', 'Template Wizard', 'The contract template editor provides a graphical interface for rapidly creating and editing smart contract templates in a modular way, allowing users to select and assemble features like building blocks to customize templates', '0.8.0', 0, 'NA', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (24, 18, 'https://g.alpha.hamsternet.io/ipfs/QmWHfgCewndaWSd2xadXTXqCNLM3kF1fN5UsrYp8ryUzV8', 'RPC Service', 'To meet developers'' needs for building various ecological projects, Hamster provides efficient and stable RPC services, empowering developers to obtain high-quality node services', '0.6.0', 0, 'https://hamsternet.io/docs/Middleware/User Guide/RPC/', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (25, 18, 'https://g.alpha.hamsternet.io/ipfs/QmZW3R6FftxoviiDc7i9Q9fVh5L6E1wFzdHConz59LooQf', 'Node Service', 'Deploy independent blockchain nodes to meet developers'' needs for accessing and using nodes', '0.10.0', 0, 'NA', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (26, 18, 'https://g.alpha.hamsternet.io/ipfs/QmbcDRJ1cqTB1oYpUwbCnpDH6AscDAaE6sRq885mvWSsSS', 'Oracle Service', 'To enable developers to quickly integrate suitable oracle solutions with no or low code, Hamster platform aggregates middleware for various oracle solutions in the market', '0.6.1', 0, 'https://hamsternet.io/docs/Middleware/User Guide/Hamslink/', '2023-11-06 10:55:43');
INSERT INTO media.t_navbar_content (id, navbar_id, icon, title, content, version, new_flag, path, create_time) VALUES (27, 18, 'https://g.alpha.hamsternet.io/ipfs/QmPPgT6WngLeZMATVNeCuuJJTfU4f2c4VMQYF8zAMhPQSX', 'Miwaspace', 'Miwaspace is a middleware marketplace that aggregates common middleware solutions in the market to meet your needs in all aspects.', '0.10.0', 0, 'https://hamsternet.io/middleware', '2023-11-06 10:55:43');

