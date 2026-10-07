INSERT INTO pages (title, value) VALUES 
    ('Главная', 'home')
;

INSERT INTO widgets (mf_name, component_name, entry_url) VALUES
    ('mf-header',    'Header',   'http://localhost:5001/remoteEntry.js'),
    ('mf-footer',    'Footer',   'http://localhost:5002/remoteEntry.js'),
    ('mf-items-list', 'ItemsList', 'http://localhost:5003/remoteEntry.js'),
    ('mf-advertising-widget', 'AdvertisingWidget', 'http://localhost:5005/remoteEntry.js'),
    ('mf-advertising-widget', 'DiscountWidget', 'http://localhost:5005/remoteEntry.js')
;