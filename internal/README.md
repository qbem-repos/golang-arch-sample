# /internal

Aplicação privada e código de bibliotecas. Este é o código que você não quer que outras pessoas importem em suas aplicações ou bibliotecas. Observe que esse padrão de layout é imposto pelo próprio compilador Go. Veja o Go 1.4 release notes para mais detalhes. Observe que você não está limitado ao diretório internal de nível superior. Você pode ter mais de um diretório internal em qualquer nível da árvore do seu projeto.

Opcionalmente, você pode adicionar um pouco de estrutura extra aos seus pacotes internos para separar o seu código interno compartilhado e não compartilhado. Não é obrigatório (especialmente para projetos menores), mas é bom ter dicas visuais que mostram o uso pretendido do pacote. Seu atual código da aplicação pode ir para o diretório ```/internal/app``` (ex. ```/internal/app/myapp```) e o código compartilhado por essas aplicações no diretório ```/internal/pkg``` (ex. ```/internal/pkg/myprivlib```).

