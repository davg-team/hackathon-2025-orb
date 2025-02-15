declare module "*.css" {
  interface IClassNames {
    className: string;
  }

  classNames = IClassNames;
  export = classNames;
}
