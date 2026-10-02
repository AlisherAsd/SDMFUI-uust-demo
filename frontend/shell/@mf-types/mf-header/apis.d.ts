
    export type RemoteKeys = 'mf-header/Header';
    type PackageType<T> = T extends 'mf-header/Header' ? typeof import('mf-header/Header') :any;