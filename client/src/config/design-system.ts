interface IAppSystem {
  APP_NAME: string;
}

interface ISurfaceSystem {
  MatFormFieldAppearance: 'fill' | 'outline';
  MatCardAppearance: 'outlined' | 'raised' | 'filled';
}

export interface IDesignSystem {
  app: IAppSystem;
  surface: ISurfaceSystem;
}
