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

export interface IDesignSystem {
  app: IAppSystem;
  typography: ITypographySystem;
  surface: ISurfaceSystem;
}
