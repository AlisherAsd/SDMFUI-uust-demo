
    export type RemoteKeys = 'mf-footer/Footer';
    type PackageType<T> = T extends 'mf-footer/Footer' ? typeof import('mf-footer/Footer') :any;