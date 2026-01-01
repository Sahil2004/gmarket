interface ISurfaceSystem {
  MatFormFieldAppearance: 'fill' | 'outline';
  MatCardAppearance: 'outlined' | 'raised' | 'filled';
}

export interface IDesignSystem {
  surface: ISurfaceSystem;
}
