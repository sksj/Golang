#include<stdio.h>
#include<string.h>
int hang[100009];
int lie[100009];
int map1[100009];
int map2[100009];
long long min(long long a,long long b)
{
    return a<b?a:b;
}
int main()
{
	int t;
	scanf("%d",&t);
	while(t--)
	{
		memset(hang,0,sizeof(hang));
		memset(lie,0,sizeof(lie));
		int n,m;
		scanf("%d %d",&n,&m);
		for(int i=1;i<=m;i++)
		{
			int a,b;
			scanf("%d %d",&a,&b);
			if(a==b)
			{
				continue;
			}
			hang[a]++;
			lie[b]++;
			map1[a]=b;
			map2[b]=a;
		}
		long long ans=0;
		for(int i=1;i<=n;i++)
		{
			if(hang[i]==1&&lie[i]==1&&(hang[map1[i]]==1&&lie[map2[i]]==1))
			{
				ans+=3;
				lie[map1[i]]=0;
				hang[map2[i]]=0;
				lie[i]=0;
				hang[i]=0;
			}
			else if(hang[i]==1&&lie[i]==1&&(hang[map1[i]]==1||lie[map2[i]]==1))
			{
				ans+=2;
				lie[map1[i]]=0;
				hang[map2[i]]=0;
				lie[i]=0;
				hang[i]=0;
			}
			else if(hang[i]==1||lie[i]==1)
			{
				ans+=1;
				if(hang[i]==1)
				{
					hang[i]=0;
					lie[map1[i]]=0;
				}
				if(lie[i]==1)
				{
					lie[i]=0;
					hang[map2[i]]=0;
				}
			}
		}
		printf("%lld\n",ans);
	}
}