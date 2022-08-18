-- +goose Up
-- +goose StatementBegin
CREATE schema europe;
CREATE TABLE IF NOT EXISTS europe.general
(
    country   VARCHAR(20) UNIQUE NOT NULL ,
    capital    VARCHAR(20) NOT NULL,
    square INTEGER  DEFAULT 0 ,
    population INTEGER  DEFAULT 0   
 ) ;
CREATE TABLE IF NOT EXISTS europe.economics
(
    country VARCHAR(20) UNIQUE  NOT NULL REFERENCES europe.general(country),
    GDP     INTEGER   NOT NULL 
    
    );
INSERT INTO europe.general VALUES 
('Albania', 'Tirana', 75000, 4000000),
('Andorra', 'Andorra la Vella', 600, 25000),
('Austria','Vienna',0 , 0),
('Belarus','Minsk', 0,0 ),
('Belgium','Brussels',0 ,0 ),
('Bosnia','Sarajevo', 0,0 ),
('Bulgaria','Sofia',0 , 0),
('Croatia','Zagreb', 0,0 ),
('Czechia','Prague', 0,0 ),
('Denmark','Copenhagen',0 ,0 ),
('Estonia','Tallinn',0 ,0 ),
('Finland','Helsinki', 0, 0),
('France','Paris', 0, 0),
('Germany','Berlin', 0,0 ),
('Greece','Athens', 0,0 ),
('Hungary','Budapest', 0,0 ),
('Iceland','Reykjavik', 0,0 ),
('Ireland','Dublin', 0,0 ),
('Italy','Rome', 0, 0),
('Latvia','Riga', 0, 0),
('Liechtenstein','Vaduz',0,0),
('Lithuania','Vilnius',0,0),
('Luxembourg','Luxembourg',0,0),
('Malta','Valletta',0,0),
('Moldova','Chisinau',0,0),
('Monaco','Monaco',0,0),
('Montenegro','Podgorica',0,0),
('Netherlands','Amsterdam',0,0),
('North Macedonia','Skopje',0,0),
('Norway','Oslo',0,0),
('Poland','Warsaw',0,0),
('Portugal','Lisbon',0,0),
('Romania','Bucharest',0,0),
('Russia','Moscow',0,0),
('San Marino','San Marino',0,0),
('Serbia','Belgrade',0,0),
('Slovakia','Bratislava',0,0),
('Slovenia','Ljubljana',0,0),
('Spain','Madrid',0,0),
('Sweden','Stockholm',0,0),
('Switzerland','Bern',0,0),
('Ukraine','Kiev',0,0),
('United Kingdom','London',0,0);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS europe.general CASCADE;
DROP TABLE IF EXISTS europe.economics CASCADE;
Drop schema IF EXISTS europe;
-- +goose StatementEnd
