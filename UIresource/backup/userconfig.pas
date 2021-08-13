unit UserConfig;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs;

type

  { TUserConfig }

  TUserConfig = class(TForm)
    procedure FormShow(Sender: TObject);
  private

  public

  end;

var
  UserConfig: TUserConfig;

implementation

{$R *.lfm}

{ TUserConfig }

procedure TUserConfig.FormShow(Sender: TObject);
begin

end;

end.

