interface IAppSystem {
  APP_NAME: string;
}

interface ITypographySystem {
  subtitleTextWeight: number;
}

interface ISurfaceSystem {
  MatFormFieldAppearance: 'fill' | 'outline';
  MatCardAppearance: 'outlined' | 'raised' | 'filled';
}

interface IColorSystem {
  theme: 'light' | 'dark' | 'system';
}

export interface IDesignSystem {
  app: IAppSystem;
  color: IColorSystem;
  typography: ITypographySystem;
  surface: ISurfaceSystem;
}
