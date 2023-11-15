CREATE TABLE IF NOT EXISTS t_faucet (
    id              int           NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name            varchar(100)  NOT NULL                           comment 'name',
    icon            varchar(200)      NULL                           comment 'icon',
    address         varchar(200)      not null                       comment 'address',
    create_time     timestamp         NULL DEFAULT CURRENT_TIMESTAMP comment 'create_time'
);

INSERT INTO t_faucet (id, name, icon, address, create_time) VALUES (1, 'Scroll Faucet','https://g.alpha.hamsternet.io/ipfs/QmXnoAtzz8FbKm5KQzo23K2WchF1Bwj2u24jLwQzh84zub','https://scroll.l2scan.co/faucet', '2023-11-15 14:37:31');
INSERT INTO t_faucet (id, name, icon, address, create_time) VALUES (2, 'BNB Chain Faucet','https://g.alpha.hamsternet.io/ipfs/Qmcv8WNHsnstNMZeFtNAYMsqTpEGgGmUfAmiJX5jCZ8qPL','https://testnet.bnbchain.org/faucet-smart', '2023-11-15 14:37:31');
INSERT INTO t_faucet (id, name, icon, address, create_time) VALUES (3, 'Solana Faucet','https://g.alpha.hamsternet.io/ipfs/QmR4nPtt1a29Xy8NWwev49cXK2dr99t3yoN21jSykoTKLs','https://faucet.solana.com', '2023-11-15 14:37:31');